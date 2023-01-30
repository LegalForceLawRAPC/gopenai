package errors

type Error struct {
	Error error
	Code  string
}

type OpenAiErrorResponse struct {
	Error struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Param   interface{} `json:"param"`
		Type    string      `json:"type"`
	} `json:"error"`
}
