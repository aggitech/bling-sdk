package bling

type ResponseErrorModel struct {
	Errors []struct {
		Error ErrorModel `json:"erro,omitempty"`
	} `json:"erros,omitempty"`
}

type ErrorModel struct {
	Code    int    `json:"cod"`
	Message string `json:"msg"`
}
