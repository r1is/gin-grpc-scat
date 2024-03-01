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
