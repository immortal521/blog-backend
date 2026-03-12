// Package response
package response

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Page[T any] struct {
	Total uint `json:"total"`
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
	Data  []T  `json:"data"`
}

const CodeSuccess = 0

func Success[T any](data T) *Response[T] {
	return &Response[T]{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	}
}

func SuccessWithMsg[T any](msg string, data T) *Response[T] {
	return &Response[T]{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	}
}

func Error(code int, msg string) *ErrorResponse {
	return &ErrorResponse{
		Code: code,
		Msg:  msg,
	}
}
