package dto

import "github.com/ethereum/go-ethereum/common"

type DeployContractResponse struct {
	Address string      `json:"address"`
	Hash    common.Hash `json:"hash"`
}

type ReadByteContractResponse struct {
	Bytecode string `json:"bytecode"`
}
