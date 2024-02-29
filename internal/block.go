package internal

import (
	"context"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

// block_number를 받지 않으면, 가장 최신 블록을 가져온다.
func (s *Server) GetBlockByNumber(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBlockHeaderByNumberRequest)
	var blockNumber *big.Int
	//var err error
	if req.BlockNumber == 0 {
		header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		blockNumber = header.Number
	} else {
		blockNumber = big.NewInt(int64(req.BlockNumber))
	}

	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, dto.ConvertBlockToResponse(*block))
}
