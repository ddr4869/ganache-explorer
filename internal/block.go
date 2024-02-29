package internal

import (
	"context"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetLatestBlockNumber(c *gin.Context) {
	header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockNumber: %v\n", header.Number)
	BlockNumber := header.Number
	c.JSON(200, gin.H{
		"block_number": BlockNumber,
	})
}

func (s *Server) GetBlockByNumber(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBlockHeaderByNumberRequest)
	BlockNumber := req.BlockNumber
	block, err := s.config.Client.BlockByNumber(context.Background(), big.NewInt(int64(BlockNumber)))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, dto.GetBlockHeaderByNumberResponse{
		BlockNumber: BlockNumber,
		Block: dto.Block{
			Number:      block.Number().Uint64(),
			GasLimit:    block.GasLimit(),
			GasUsed:     block.GasUsed(),
			Time:        block.Time(),
			Difficulty:  block.Difficulty().Uint64(),
			NumberU64:   block.NumberU64(),
			MixDigest:   block.MixDigest().Hex(),
			Nonce:       block.Nonce(),
			Coinbase:    block.Coinbase().Hex(),
			Root:        block.Root().Hex(),
			ParentHash:  block.ParentHash().Hex(),
			TxHash:      block.TxHash().Hex(),
			ReceiptHash: block.ReceiptHash().Hex(),
			UncleHash:   block.UncleHash().Hex(),

			TransactionCount: len(block.Transactions()),
		},
	})
}
