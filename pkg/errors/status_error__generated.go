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
	case NotFound:
		return "NotFound"
	case Conflict:
		return "Conflict"
	case InternalServerError:
		return "InternalServerError"
	case DatabaseInternalServerError:
		return "DatabaseInternalServerError"
	case GenerateTokenError:
		return "GenerateTokenError"
	case UploadStorage:
		return "UploadStorage"
	case GetDownloadLink:
		return "GetDownloadLink"
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
	case NotFound:
		return "NotFound"
	case Conflict:
		return "Conflict"
	case InternalServerError:
		return "InternalServerError"
	case DatabaseInternalServerError:
		return "查询数据库失败"
	case GenerateTokenError:
		return "Token生成失败"
	case UploadStorage:
		return "存储出错"
	case GetDownloadLink:
		return "获取下载链接出错"
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
	case NotFound:
		return true
	case Conflict:
		return true
	case InternalServerError:
		return true
	case DatabaseInternalServerError:
		return true
	case GenerateTokenError:
		return true
	case UploadStorage:
		return true
	case GetDownloadLink:
		return true
	}
	return false
}
