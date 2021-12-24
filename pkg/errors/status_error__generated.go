package errors

import (
	github_com_go_courier_statuserror "github.com/go-courier/statuserror"
)

var _ interface {
	github_com_go_courier_statuserror.StatusError
} = (*StatusError)(nil)

func (v StatusError) StatusErr() *github_com_go_courier_statuserror.StatusErr {
	return &github_com_go_courier_statuserror.StatusErr{
		Key:            v.Key(),
		Code:           v.Code(),
		Msg:            v.Msg(),
		CanBeTalkError: v.CanBeTalkError(),
	}
}

func (v StatusError) Unwrap() error {
	return v.StatusErr()
}

func (v StatusError) Error() string {
	return v.StatusErr().Error()
}

func (v StatusError) StatusCode() int {
	return github_com_go_courier_statuserror.StatusCodeFromCode(int(v))
}

func (v StatusError) Code() int {
	if withServiceCode, ok := (interface{})(v).(github_com_go_courier_statuserror.StatusErrorWithServiceCode); ok {
		return withServiceCode.ServiceCode() + int(v)
	}
	return int(v)

}

func (v StatusError) Key() string {
	switch v {
	case Unauthorized:
		return "Unauthorized"
	case InvalidToken:
		return "InvalidToken"
	case TokenExpired:
		return "TokenExpired"
	case TokenNotFoundOrExpired:
		return "TokenNotFoundOrExpired"
	case UnmatchedPassword:
		return "UnmatchedPassword"
	case Forbidden:
		return "Forbidden"
	case FirmwareReleased:
		return "FirmwareReleased"
	case InternalServerError:
		return "InternalServerError"
	case DatabaseInternalServerError:
		return "DatabaseInternalServerError"
	case GenerateTokenError:
		return "GenerateTokenError"
	}
	return "UNKNOWN"
}

func (v StatusError) Msg() string {
	switch v {
	case Unauthorized:
		return "Unauthorized"
	case InvalidToken:
		return "登录状态已失效"
	case TokenExpired:
		return "登录状态已失效"
	case TokenNotFoundOrExpired:
		return "登录状态已失效"
	case UnmatchedPassword:
		return "密码错误"
	case Forbidden:
		return "Forbidden"
	case FirmwareReleased:
		return "固件多次发布"
	case InternalServerError:
		return "InternalServerError"
	case DatabaseInternalServerError:
		return "查询数据库失败"
	case GenerateTokenError:
		return "Token生成失败"
	}
	return "-"
}

func (v StatusError) CanBeTalkError() bool {
	switch v {
	case Unauthorized:
		return true
	case InvalidToken:
		return true
	case TokenExpired:
		return true
	case TokenNotFoundOrExpired:
		return true
	case UnmatchedPassword:
		return true
	case Forbidden:
		return true
	case FirmwareReleased:
		return true
	case InternalServerError:
		return false
	case DatabaseInternalServerError:
		return true
	case GenerateTokenError:
		return true
	}
	return false
}
