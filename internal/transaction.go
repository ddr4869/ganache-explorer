package internal

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransactions(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionRequest)
	blockNumber := int64(req.BlockNumber)

	block, err := s.config.Client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, dto.GetTransactionResponse{
		Transactions: block.Transactions(),
		Tx_length:    len(block.Transactions()),
	})

	// })
	// result := make([]types.Receipt, 0)
	// for i, tx := range block.Transactions() {
	// 	if i > 1 {
	// 		break
	// 	}
	// 	chainID, err := s.config.Client.NetworkID(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
	// 		fmt.Println(from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
	// 	}
	// 	receipt, err := s.config.Client.TransactionReceipt(context.Background(), tx.Hash())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("Receipt: %v\n", receipt)
	// 	result = append(result, *receipt)
	// }

	// c.JSON(200, gin.H{
	// 	"receipt": result,
	// })
}

func (s *Server) GetTransactionReceipt(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionReceiptRequest)
	blockNumber := int64(req.BlockNumber)

	block, err := s.config.Client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
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
		log.Printf("tx hash: %v\n", tx.Hash())
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
