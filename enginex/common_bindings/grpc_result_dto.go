package common_bindings

type GRPCResultDTO struct {
	OK           bool   `json:"ok"`
	MsgErrorCode string `json:"msg_error_code"`
	MsgError     string `json:"msg_error"`
}
