package dto

type GetTransactionRequest struct {
	BlockNumber int `form:"block_number"`
}

type GetTransactionHashesRequest struct {
	BlockNumber int `form:"block_number"`
}

type GetTransactionReceiptRequest struct {
	BlockNumber int    `form:"block_number" required:"true" binding:"required"`
	Hash        string `form:"tx_hash"`
}

type GetTransactionByHashRequest struct {
	BlockNumber int    `form:"block_number" required:"true" binding:"required"`
	Hash        string `form:"tx_hash"`
}
