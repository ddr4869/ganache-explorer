package internal

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ddr4869/ether-go/internal/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func (s *Server) Transfer(c *gin.Context) {
	req := c.MustGet("req").(dto.TransferRequest)
	privateKey, err := crypto.HexToECDSA(utils.Remove0xPrefix(req.From_private))
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind private key")
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Cannot cast public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get nonce")
		return
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

	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get chainId")
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to sign tx")
		return
	}

	err = s.config.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to send tx")
		return
	}
	c.JSON(200, dto.TransactionHashResponse{
		Hash: signedTx.Hash(),
	})
}

func (s *Server) TransferToken(c *gin.Context) {
	req := c.MustGet("req").(dto.TransferTokenRequest)
	privateKey, err := crypto.HexToECDSA(utils.Remove0xPrefix(req.From_private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(req.To_address)
	tokenAddress := common.HexToAddress("0x3b90CD2eF2F65F569B59B32e1a4939D1c91DF55e")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.NewKeccakState()
	hash.Write(transferFnSignature)
	log.Printf("hash 1: %x", hash.Sum(nil)[:4]) // 0xa9059cbb
	methodID := hash.Sum(nil)[:4]
	fmt.Println("method: ", hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString(req.Amount, 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

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

	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = s.config.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	c.JSON(200, dto.TransactionHashResponse{
		Hash: signedTx.Hash(),
	})
}
