package apis

type Resp struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

func SuccessResp() Resp {
	return Resp{
		StatusCode:    0,
		StatusMessage: "",
	}
}

// 上传代码成功
func SuccessRespOK() Resp {
	return Resp{
		StatusCode:    0,
		StatusMessage: "ok",
	}
}

// 审计成功
func SuccessRespWithMsg(msg string) Resp {
	return Resp{
		StatusCode:    0,
		StatusMessage: msg,
	}
}

// 失败的返回
func FailResp(faildMsg string) Resp {
	return Resp{
		StatusCode:    -1,
		StatusMessage: faildMsg,
	}
}

// 正在审计中
func AuditingResp() Resp {
	return Resp{
		StatusCode:    1,
		StatusMessage: "正在审计中",
	}
}

// 完成审计
func FinishAuditResp() Resp {
	return Resp{
		StatusCode:    0,
		StatusMessage: "审计完成",
	}
}

// 完成审计，但是审计失败
func FinishAuditFailResp(faildMsg string) Resp {
	return Resp{
		StatusCode:    2,
		StatusMessage: faildMsg,
	}
}
