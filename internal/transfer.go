package internal

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ddr4869/ether-go/internal/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/gin-gonic/gin"
)

func (s *Server) Transfer(c *gin.Context) {
	req := c.MustGet("req").(dto.TransferRequest)
	nonce, privateKey, err := s.GetNonceAndPKfromKeyString(req.From_private)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get nonce")
	}

	value := big.NewInt(req.Value)    // in wei (1 eth)
	gasLimit := uint64(req.Gas_limit) // in units
	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to guess gas price")
		return
	}

	toAddress := common.HexToAddress(req.To_address)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := s.SignAndSendTransaction(tx, privateKey)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to SignAndSendTransaction")
		return
	}

	c.JSON(200, dto.TransactionHashResponse{
		Hash: signedTx.Hash(),
	})
}

func (s *Server) TransferToken(c *gin.Context) {
	req := c.MustGet("req").(dto.TransferTokenRequest)
	nonce, privateKey, err := s.GetNonceAndPKfromKeyString(req.From_private)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get nonce")
	}

	value := big.NewInt(0)
	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(req.To_address)
	tokenAddress := common.HexToAddress(req.Token_address) // 0x3b90CD2eF2F65F569B59B32e1a4939D1c91DF55e

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.NewKeccakState()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString(req.Amount, 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := s.config.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	signedTx, err := s.SignAndSendTransaction(tx, privateKey)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to SignAndSendTransaction")
		return
	}

	c.JSON(200, dto.TransactionHashResponse{
		Hash: signedTx.Hash(),
	})
}

func (s *Server) SignAndSendTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = s.config.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, err
}

func (s *Server) SignRawTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (string, error) {
	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	return rawTxHex, err
}

func (s *Server) GetNonceAndPKfromKeyString(key string) (uint64, *ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(utils.Remove0xPrefix(key))
	if err != nil {
		return 0, nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return 0, nil, err
	}
	return nonce, privateKey, nil
}
