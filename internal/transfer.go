package internal

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ddr4869/ether-go/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func (s *Server) Transfer(c *gin.Context) {
	req := c.MustGet("req").(dto.TransferRequest)
	privateKey, err := crypto.HexToECDSA(utils.Remove0xPrefix(req.From_private))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(req.Value)    // in wei (1 eth)
	gasLimit := uint64(req.Gas_limit) // in units
	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(req.To_address)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = s.config.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"txHash": signedTx.Hash().Hex(),
	})
}
