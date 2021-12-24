package errors

import "net/http"

//go:generate tools gen error StatusError
type StatusError int

func (StatusError) ServiceCode() int {
	return 999 * 1e3
}

const (
	// InternalServerError
	InternalServerError StatusError = http.StatusInternalServerError*1e6 + iota + 1
	// @errTalk 查询数据库失败
	DatabaseInternalServerError
	// @errTalk Token生成失败
	GenerateTokenError
)

const (
	// @errTalk Unauthorized
	Unauthorized StatusError = http.StatusUnauthorized*1e6 + iota + 1
	// @errTalk 登录状态已失效
	InvalidToken
	// @errTalk 登录状态已失效
	TokenExpired
	// @errTalk 登录状态已失效
	TokenNotFoundOrExpired
	// @errTalk 密码错误
	UnmatchedPassword
)

const (
	// @errTalk Forbidden
	Forbidden StatusError = http.StatusForbidden*1e6 + iota + 1
	// @errTalk 固件多次发布
	FirmwareReleased
)

const (
	// @errTalk NotFound
	NotFound = http.StatusNotFound*1e6 + iota + 1
)
