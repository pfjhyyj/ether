package define

type GetSMSCodeRequest struct {
	Phone string `json:"phone"`
}

type VerifySMSCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
