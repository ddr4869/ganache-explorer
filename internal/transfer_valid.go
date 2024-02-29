package internal

import (
	"log"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) TransferValid(c *gin.Context) {
	var req dto.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Fatalf("Failed to bind request: %v", err)
		return
	}
	c.Set("req", req)
	c.Next()
}
