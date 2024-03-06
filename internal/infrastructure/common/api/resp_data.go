package api

// RespData 响应
type RespData[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// Page 分页
type Page[T any] struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}

// Success 成功响应
func Success[T any](data T, message string) RespData[T] {
	respData := RespData[T]{
		Code:    200,
		Message: message,
		Data:    data,
	}
	return respData
}

// Fail 失败响应
func Fail[T any](message string) RespData[T] {
	var zeroValue T
	respData := RespData[T]{
		Code:    500,
		Message: message,
		Data:    zeroValue,
	}
	return respData
}

// UnAuth 为鉴权响应
func UnAuth[T any](message string) RespData[T] {
	var zeroValue T
	respData := RespData[T]{
		Code:    401,
		Message: message,
		Data:    zeroValue,
	}
	return respData
}
