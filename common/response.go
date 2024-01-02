package common

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Page struct {
	Current  int64         `json:"current"`
	PageSize int64         `json:"pageSize"`
	Total    int64         `json:"total"`
	List     []interface{} `json:"list"`
}

type PageResponse struct {
	Response
	Data Page `json:"data"`
}
