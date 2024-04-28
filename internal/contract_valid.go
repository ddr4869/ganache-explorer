package internal

import (
	"net/http"

	"github.com/ddr4869/ganache-explorer/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) DeployContractValid(c *gin.Context) {
	var req dto.DeployContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) WriteContractValid(c *gin.Context) {
	var req dto.WriteContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) ReadByteContractValid(c *gin.Context) {
	var req dto.ReadByteContractRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) LoadStoreContractValid(c *gin.Context) {
	var req dto.LoadStoreContractRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) LoadErc20ContractValid(c *gin.Context) {
	var req dto.LoadStoreContractRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
