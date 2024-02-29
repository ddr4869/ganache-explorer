package internal

import (
	"log"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransactionValid(c *gin.Context) {
	var req dto.GetTransactionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Fatalf("Failed to bind request: %v", err)
		return
	}
	c.Set("req", req)
	c.Next()
}
