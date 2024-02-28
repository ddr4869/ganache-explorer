package internal

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetBlockHeaderByNumber(c *gin.Context) {
	header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockNumber: %v\n", header.Number)
	BlockNumber := header.Number

	block, err := s.config.Client.BlockByNumber(context.Background(), BlockNumber)
	if err != nil {
		log.Fatal(err)
	}
	count, err := s.config.Client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"blockHeader":       header,
		"block":             block,
		"transaction_count": count,
	})
}
