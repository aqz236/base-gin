package errors

import "fmt"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Detail)
}

func NewAppError(code, message, detail string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// 预定义错误类型
var (
	ErrUserNotFound    = NewAppError("USER_NOT_FOUND", "用户不存在", "")
	ErrEmailExists     = NewAppError("EMAIL_EXISTS", "邮箱已存在", "")
	ErrInvalidEmail    = NewAppError("INVALID_EMAIL", "邮箱格式不正确", "")
	ErrInvalidName     = NewAppError("INVALID_NAME", "用户名格式不正确", "")
	ErrInvalidPassword = NewAppError("INVALID_PASSWORD", "密码格式不正确", "")
	ErrInternalServer  = NewAppError("INTERNAL_ERROR", "服务器内部错误", "")
	ErrBadRequest      = NewAppError("BAD_REQUEST", "请求参数错误", "")
)
