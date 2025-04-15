package model

// Response 统一响应结构
type Response struct {
	Status  int         `json:"status"`  // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 响应数据
}

// Success 成功响应
func Success(data interface{}, message string) *Response {
	response := &Response{
		Status:  200,
		Message: message,
		Data:    data,
	}
	return response
}

// Error 错误响应
func Error(status int, message string) *Response {
	response := &Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	return response
}
