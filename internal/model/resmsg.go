package model

// ResMsg 响应json信息的结构体
type ResMsg struct {
	Msg  string      `json:"message"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}
