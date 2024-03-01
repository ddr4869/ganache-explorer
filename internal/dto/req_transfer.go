package dto

type TransferRequest struct {
	//From_address string `json:"from_address"`
	From_private string `json:"from_private" binding:"required"`
	To_address   string `json:"to_address" binding:"required"`
	Value        int64  `json:"value" binding:"required"`
	Gas_limit    int64  `json:"gas_limit" binding:"required"`
}

type TransferTokenRequest struct {
	//From_address string `json:"from_address"`
	From_private  string `json:"from_private" binding:"required"`
	To_address    string `json:"to_address" binding:"required"`
	Token_address string `json:"token_address"`
	Amount        string `json:"amount" binding:"required"`
}
