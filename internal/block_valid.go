package internal

import (
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetBlockByNumberValid(c *gin.Context) {
	var req dto.GetBlockHeaderByNumberRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
