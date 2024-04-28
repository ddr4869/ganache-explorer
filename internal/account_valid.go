package internal

import (
	"net/http"

	"github.com/ddr4869/ganache-explorer/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetBalanceValid(c *gin.Context) {
	var req dto.GetBalanceRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}
	c.Set("req", req)
	c.Next()
}
