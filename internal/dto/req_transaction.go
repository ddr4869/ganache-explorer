package dto

type GetTransactionRequest struct {
	BlockNumber int `form:"block_number"`
}

type GetTransactionReceiptRequest struct {
	BlockNumber int `form:"block_number"`
}
