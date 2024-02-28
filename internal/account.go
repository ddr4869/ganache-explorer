package internal

import (
	"context"
	"log"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetBalance(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBalanceRequest)
	log.Printf("Address: %s\n", req.Address)
	account := common.HexToAddress(req.Address)
	balance, err := s.config.Ganache.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance: %s\n", balance)
	c.JSON(200, gin.H{
		"balance": balance,
	})
}
