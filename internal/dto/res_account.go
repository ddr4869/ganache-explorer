package dto

import "math/big"

type GetBalanceResponse struct {
	Wei   *big.Int   `json:"wei"`
	Ether *big.Float `json:"ether"`
}
