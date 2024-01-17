package common

type SystemError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error
}

func (e *SystemError) Error() string {
	return e.Message
}
