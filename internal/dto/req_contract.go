package dto

type DeployContractRequest struct {
	From_private string `json:"from_private"`
	To_address   string `json:"to_address"`
	Contract     string `json:"contract"`
}

type WriteContractRequest struct {
	From_private     string `json:"from_private"`
	Contract_address string `json:"contract_address"`
	Key              string `json:"key"`
	Value            string `json:"value"`
}

type ReadByteContractRequest struct {
	Contract_address string `json:"contract_address"`
}

type LoadStoreContractRequest struct {
	Contract_address string `json:"contract_address"`
	Key              string `json:"key"`
}
