package dto

type GetBalanceRequest struct {
	Address     string `form:"address"`
	BlockNumber int    `form:"block_number"`
}
