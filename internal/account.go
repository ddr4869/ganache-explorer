package internal

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ddr4869/ganache-explorer/internal/dto"
	"github.com/ddr4869/ganache-explorer/internal/utils"
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
