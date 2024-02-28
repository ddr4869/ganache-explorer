package internal

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransaction(c *gin.Context) {
	blockNumber := big.NewInt(5379777)
	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for i, tx := range block.Transactions() {
		if i > 10 {
			break
		}
		log.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		log.Println(tx.Value().String())    // 10000000000000000
		log.Println(tx.Gas())               // 105000
		log.Println(tx.GasPrice().Uint64()) // 102000000000
		log.Println(tx.Nonce())             // 110644
		log.Println(tx.Data())              // []
		log.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := s.config.Client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println(from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
		receipt, err := s.config.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	c.JSON(200, gin.H{
		"tx": block.Transactions(),
	})
}
