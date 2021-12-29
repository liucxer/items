package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/statuserror"
)

//go:generate tools gen error StatusError
type StatusError int

func (StatusError) ServiceCode() int {
	return 999 * 1e3
}

func (v StatusError) WithDes(des ...interface{}) *statuserror.StatusErr {
	ret := v.StatusErr()
	if len(des) == 0 {
		return ret
	}
	return ret.WithDesc(fmt.Sprint(des))
}

func (v StatusError) WithMsg(msg ...interface{}) *statuserror.StatusErr {
	ret := v.StatusErr()
	if len(msg) == 0 {
		return ret
	}
	ret.CanBeTalkError = true
	return ret.WithMsg(fmt.Sprint(msg))
}

const (
	// @errTalk InternalServerError
	InternalServerError StatusError = http.StatusInternalServerError*1e6 + iota + 1
	// @errTalk 查询数据库失败
	DatabaseInternalServerError
	// @errTalk Token生成失败
	GenerateTokenError
	// @errTalk 存储出错
	UploadStorage
	// @errTalk 获取下载链接出错
	GetDownloadLink
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
	NotFound StatusError = http.StatusNotFound*1e6 + iota + 1
)

const (
	// @errTalk Conflict
	Conflict StatusError = http.StatusConflict*1e6 + iota + 1
)

func DBError(err error) *statuserror.StatusErr {
	if sqlx.DBErr(err).IsNotFound() {
		return NotFound.WithDes()
	} else if sqlx.DBErr(err).IsConflict() {
		return Conflict.WithDes()
	} else {
		return InternalServerError.WithDes(err)
	}
}

var New = errors.New
