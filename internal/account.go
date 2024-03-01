package internal

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ddr4869/ether-go/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetBalance(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBalanceRequest)
	account := common.HexToAddress(req.Address)
	var blockNumber *big.Int
	if req.BlockNumber != 0 {
		blockNumber = big.NewInt(int64(req.BlockNumber))
	}
	balance, err := s.config.Client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get balance")
		return
	}
	ether := utils.ConvertWeiToEther(*balance)
	c.JSON(http.StatusOK, dto.GetBalanceResponse{
		Wei:   balance,
		Ether: ether,
	})
}

// 트랜잭션을 제출하거나 확인을 기다리는 등의 보류 중인 계정 잔액을 조회합니다.
func (s *Server) GetPendingBalance(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBalanceRequest)
	account := common.HexToAddress(req.Address)
	balance, err := s.config.Client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get balance")
		return

	}
	ether := utils.ConvertWeiToEther(*balance)
	c.JSON(http.StatusOK, dto.GetBalanceResponse{
		Wei:   balance,
		Ether: ether,
	})
}
