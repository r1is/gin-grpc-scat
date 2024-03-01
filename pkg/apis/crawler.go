package apis

type GetCurrentBlockNumberReq struct {
	ID string `json:"id"`
}

type GetCurrentBlockNumberResp struct {
	BlockNumber int64 `json:"block_number"`
	Resp
}

type UpdateCurrentBlockNumberReq struct {
	ID          string `json:"id"`
	BlockNumber int64  `json:"block_number"`
}

type UpdateCurrentBlockNumberResp struct {
	Resp
}

type AddBlockQueueReq struct {
	ID          string `json:"id"`
	BlockNumber int64  `json:"block_number"`
}

type AddBlockQueueResp struct {
	Resp
}

type CollectContractsFromBlockReq struct {
	ChainName   string `json:"chain_name"`
	BlockNumber int64  `json:"block_number"`
}

type CollectContractsFromBlockResp struct {
	Resp
}

type CollectContractReq struct {
	ChainName string `json:"chain_name"`
	Address   string `json:"address"`
}

type CollectContractResp struct {
	ByteCode   string `json:"byte_code"`
	SourceCode string `json:"source_code"`
	Resp
}
