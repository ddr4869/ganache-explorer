package dto

type GetBlockHeaderByNumberRequest struct {
	BlockNumber int `form:"block_number"`
}
