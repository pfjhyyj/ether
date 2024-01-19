package common

type SystemError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error
}

func (e *SystemError) Error() string {
	return e.Msg
}
