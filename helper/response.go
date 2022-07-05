package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

func APIResponseFailed(message string, code int, success bool) *Meta {
	meta := Meta{
		Message: message,
		Code:    code,
		Success: success,
	}

	return &meta
}
