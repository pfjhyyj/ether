package common

type PageRequest struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}
