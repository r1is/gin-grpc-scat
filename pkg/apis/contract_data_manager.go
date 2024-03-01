package apis

import models "pkg/contract_models"

type GetCommonContractReq struct {
	ID string `json:"id"`
}

type GetCommonContractResp struct {
	Contract models.CommonContract `json:"contract"`
	Resp
}

type AddCommonContractReq struct {
	Contract models.CommonContract `json:"contract"`
}

type AddCommonContractResp struct {
	ID string `json:"id"`
	Resp
}

type UpdateCommonContractReq struct {
	ID       string                `json:"id"`
	Contract models.CommonContract `json:"contract"`
}

type UpdateCommonContractResp struct {
	Resp
}

type GetDeployedContractReq struct {
	ID string `json:"id"`
}

type GetDeployedContractResp struct {
	Contract models.DeployedContract `json:"contract"`
	Resp
}

type AddDeployedContractReq struct {
	Contract models.DeployedContract `json:"contract"`
}

type AddDeployedContractResp struct {
	ID string `json:"id"`
	Resp
}

type UpdateDeployedContractReq struct {
	ID       string                  `json:"id"`
	Contract models.DeployedContract `json:"contract"`
}

type UpdateDeployedContractResp struct {
	Resp
}

type GetOldestUncheckedContractIDReq struct {
}

type GetOldestUncheckedContractIDResp struct {
	ID string `json:"id"`
	Resp
}
