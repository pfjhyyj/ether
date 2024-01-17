package common

const (
	Ok           = 0
	UnknownError = 10001
	DbError      = 10002

	RequestError = 20001
	AuthError    = 20002
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Page struct {
	Current  int         `json:"current"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}
