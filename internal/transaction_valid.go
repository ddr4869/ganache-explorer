package internal

import (
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransactionByHashValid(c *gin.Context) {
	var req dto.GetTransactionByHashRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind tx hash")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetTransactionsValid(c *gin.Context) {
	var req dto.GetTransactionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetTransactionHashesValid(c *gin.Context) {
	var req dto.GetTransactionHashesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetTransactionReceiptValid(c *gin.Context) {
	var req dto.GetTransactionReceiptRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind block number or hasn")
		return
	}
	c.Set("req", req)
	c.Next()
}
