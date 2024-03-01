package internal

import (
	"context"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

// block_number를 받지 않으면, 가장 최신 블록을 가져온다.
func (s *Server) GetBlockByNumber(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBlockHeaderByNumberRequest)
	blockNumber, err := s.CheckBlockNumber(req.BlockNumber)
	if err != nil {
		return
	}
	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, dto.ConvertBlockToResponse(*block))
}

func (s *Server) CheckBlockNumber(BlockNumber int) (*big.Int, error) {
	var blockNumber *big.Int
	if BlockNumber == 0 {
		header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		blockNumber = header.Number
	} else {
		blockNumber = big.NewInt(int64(BlockNumber))
	}
	return blockNumber, nil
}

func (s *Server) SubscribeBlock(c *gin.Context) {

	headers := make(chan *types.Header) // 최신 블록 헤더를 받을 chan 생성
	sub, err := s.config.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to subscribe new head")
		return
	}

	for {
		select {
		case err := <-sub.Err():
			dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Received error from subscription")
			return
		case header := <-headers:
			log.Printf("block hash: %s", header.Hash().Hex())
			log.Printf("block number: %s", header.Number.String())
			block, err := s.config.Client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get block by hash")
				return
			}
			log.Printf("block tx len: %d", block.Transactions().Len()) // 3
		}
	}
}
