package dto

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// types.Transaction example
// {
// 	"type": "0x2",
// 	"chainId": "0xaa36a7",
// 	"nonce": "0x0",
// 	"to": "0x635d73d6c5ebdf9bae66b1fa9e9e07361305e7b9",
// 	"gas": "0x186a0",
// 	"gasPrice": null,
// 	"maxPriorityFeePerGas": "0x3b9aca000",
// 	"maxFeePerGas": "0x3f5476a18",
// 	"value": "0x0",
// 	"input": "0x095ea7b3000000000000000000000000e877139db8095dd59fcbbfd65a02ae08592ac8ea00000000000000000000000000000000000000000000000000000000004c4b40",
// 	"accessList": [],
// 	"v": "0x0",
// 	"r": "0x74c78978f7c39afa28526c57acf0ab61010467cc8a9fb65ba4b2c046ea4e6b69",
// 	"s": "0x78d5b577ac5019a76ebcb6d0cac8d8258e2bf0bc39a04122b9af1e3a70e30927",
// 	"yParity": "0x0",
// 	"hash": "0x35b32a3ff7470ed49db25b192608858677ff2a98fe31231a85664ecac98ae288"
// },

type GetTransactionResponse struct {
	Transactions types.Transactions `json:"transactions"`
	Tx_length    int                `json:"tx_length"`
}
