package internal

import (
	"net/http"

	"github.com/ddr4869/ganache-explorer/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) TransferValid(c *gin.Context) {
	var req dto.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind req")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) TransferTokenValid(c *gin.Context) {
	var req dto.TransferTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind req")
		return
	}
	c.Set("req", req)
	c.Next()
}
