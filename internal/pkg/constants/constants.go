package constants

const (
	// HTTP Status Messages
	StatusOK     = "success"
	StatusError  = "error"
	StatusFailed = "failed"

	// User related
	UserCreated  = "用户创建成功"
	UserUpdated  = "用户更新成功"
	UserDeleted  = "用户删除成功"
	UserNotFound = "用户不存在"

	// Validation Messages
	InvalidEmail = "邮箱格式不正确"
	InvalidName  = "用户名格式不正确"
	EmailExists  = "邮箱已存在"

	// Server Messages
	InternalError = "服务器内部错误"
	BadRequest    = "请求参数错误"
)
