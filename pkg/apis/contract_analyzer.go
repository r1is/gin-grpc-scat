package apis

import models "pkg/contract_models"

type AnalyzeByteCodeReq struct {
	ByteCode string `json:"byte_code"`
	ToolName string `json:"tool_name"`
}

type AnalyzeByteCodeResp struct {
	Result models.ByteCodeAnalyzeResult `json:"result"`
	Resp
}

type GetDocReq struct {
}

type GetDocResp struct {
	Doc string `json:"doc"`
	Resp
}

// RequestId  string `json:"requestId"` tools漏洞返回的结果
type AnalyzeSourceCodeResp struct {
	RequestId string                         `json:"requestId"`
	Result    models.SourceCodeAnalyzeResult `json:"result"`
	Resp
}

type AnalyzeSourceCodeReq struct {
	RequestId  string `json:"requestId"`
	SourceCode string `json:"source_code"`
	ToolName   string `json:"tool_name"`
}

// 临时满足邱老师需求
type RespOK struct {
	RequestId     string `json:"requestId"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
