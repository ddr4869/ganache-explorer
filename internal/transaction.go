package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransaction(c *gin.Context) {
	header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockNumber: %v\n", header.Number)
	blockNumber := header.Number

	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]types.Receipt, 0)
	for i, tx := range block.Transactions() {
		if i > 1 {
			break
		}
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
		log.Printf("Receipt: %v\n", receipt)
		result = append(result, *receipt)
	}

	c.JSON(200, gin.H{
		"receipt": result,
	})
}
