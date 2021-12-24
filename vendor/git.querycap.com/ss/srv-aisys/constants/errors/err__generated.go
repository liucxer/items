package errors

import (
	github_com_go_courier_statuserror "github.com/go-courier/statuserror"
)

var _ interface {
	github_com_go_courier_statuserror.StatusError
} = (*err)(nil)

func (v err) StatusErr() *github_com_go_courier_statuserror.StatusErr {
	return &github_com_go_courier_statuserror.StatusErr{
		Key:            v.Key(),
		Code:           v.Code(),
		Msg:            v.Msg(),
		CanBeTalkError: v.CanBeTalkError(),
	}
}

func (v err) Unwrap() error {
	return v.StatusErr()
}

func (v err) Error() string {
	return v.StatusErr().Error()
}

func (v err) StatusCode() int {
	return github_com_go_courier_statuserror.StatusCodeFromCode(int(v))
}

func (v err) Code() int {
	if withServiceCode, ok := (interface{})(v).(github_com_go_courier_statuserror.StatusErrorWithServiceCode); ok {
		return withServiceCode.ServiceCode() + int(v)
	}
	return int(v)

}

func (v err) Key() string {
	switch v {
	case BadRequestError:
		return "BadRequestError"
	case OTAInvalidParam:
		return "OTAInvalidParam"
	case CamElectricFenceInvalid:
		return "CamElectricFenceInvalid"
	case CamEnableSectionInvalid:
		return "CamEnableSectionInvalid"
	case CamLonLatInvalid:
		return "CamLonLatInvalid"
	case SysConfigLonLatInvalid:
		return "SysConfigLonLatInvalid"
	case SysConfigInvalid:
		return "SysConfigInvalid"
	case UnauthorizedError:
		return "UnauthorizedError"
	case UserLoginInvalidCaptcha:
		return "UserLoginInvalidCaptcha"
	case UserLoginInvalidToken:
		return "UserLoginInvalidToken"
	case UserLoginInvalidPassword:
		return "UserLoginInvalidPassword"
	case ForbiddenError:
		return "ForbiddenError"
	case OTABusy:
		return "OTABusy"
	case OTADiskLimited:
		return "OTADiskLimited"
	case OTAUnsupportedModel:
		return "OTAUnsupportedModel"
	case OTAUploadLimited:
		return "OTAUploadLimited"
	case OTAMd5Checksum:
		return "OTAMd5Checksum"
	case OTAInvalidVersion:
		return "OTAInvalidVersion"
	case OTAMaxModelInstall:
		return "OTAMaxModelInstall"
	case OTAInvalidURL:
		return "OTAInvalidURL"
	case CamSceneConfLimited:
		return "CamSceneConfLimited"
	case CamConfLimited:
		return "CamConfLimited"
	case UploadBusy:
		return "UploadBusy"
	case UploadFileSizeLimited:
		return "UploadFileSizeLimited"
	case UploadDiskLimited:
		return "UploadDiskLimited"
	case SysConfigPermissionDenied:
		return "SysConfigPermissionDenied"
	case NotFoundError:
		return "NotFoundError"
	case DBNotFound:
		return "DBNotFound"
	case CamNotExists:
		return "CamNotExists"
	case DBConflict:
		return "DBConflict"
	case InternalServerError:
		return "InternalServerError"
	case UserLoginGenTokenFailed:
		return "UserLoginGenTokenFailed"
	case UserLoginInvalidUsername:
		return "UserLoginInvalidUsername"
	case DBInternal:
		return "DBInternal"
	case EventResReadFailed:
		return "EventResReadFailed"
	case OTALicenseValidate:
		return "OTALicenseValidate"
	case OTAHandleUpload:
		return "OTAHandleUpload"
	case OTAFetchRemoteFirmware:
		return "OTAFetchRemoteFirmware"
	case OTASubmitFailed:
		return "OTASubmitFailed"
	case OTAPadModelInfoFailed:
		return "OTAPadModelInfoFailed"
	case OTARemoveModelFailed:
		return "OTARemoveModelFailed"
	case OTAUpgradeFailed:
		return "OTAUpgradeFailed"
	case CamAnalyzeReqError:
		return "CamAnalyzeReqError"
	case LicenseValidateByAnalysis:
		return "LicenseValidateByAnalysis"
	case LicenseInstallFailed:
		return "LicenseInstallFailed"
	case LicenseFileNotExist:
		return "LicenseFileNotExist"
	case LicenseFileHashFailed:
		return "LicenseFileHashFailed"
	case UploadHandleFailed:
		return "UploadHandleFailed"
	}
	return "UNKNOWN"
}

func (v err) Msg() string {
	switch v {
	case BadRequestError:
		return ""
	case OTAInvalidParam:
		return "升级请求参数非法"
	case CamElectricFenceInvalid:
		return "电子围栏参数错误"
	case CamEnableSectionInvalid:
		return "场景分析时段数据错误"
	case CamLonLatInvalid:
		return "相机经纬度数据错误 纬度:[-90,90];经度:[-180,180]"
	case SysConfigLonLatInvalid:
		return "设备经纬度数据错误 纬度:[-90,90];经度:[-180,180]"
	case SysConfigInvalid:
		return "配置项不合法"
	case UnauthorizedError:
		return ""
	case UserLoginInvalidCaptcha:
		return "验证码错误"
	case UserLoginInvalidToken:
		return "登录已过期"
	case UserLoginInvalidPassword:
		return "密码错误"
	case ForbiddenError:
		return "Forbidden"
	case OTABusy:
		return "升级任务忙"
	case OTADiskLimited:
		return "磁盘不足"
	case OTAUnsupportedModel:
		return "证书不支持该模型"
	case OTAUploadLimited:
		return "上传文件过大"
	case OTAMd5Checksum:
		return "md5校验失败"
	case OTAInvalidVersion:
		return "低于当前版本"
	case OTAMaxModelInstall:
		return "可安装达到上限"
	case OTAInvalidURL:
		return "无效链接"
	case CamSceneConfLimited:
		return "场景配置达到上限"
	case CamConfLimited:
		return "摄像机配置达到上限"
	case UploadBusy:
		return "上传任务繁忙"
	case UploadFileSizeLimited:
		return "上传文件大小限制"
	case UploadDiskLimited:
		return "上传任务磁盘不足"
	case SysConfigPermissionDenied:
		return "无系统配置权限"
	case NotFoundError:
		return "NotFound"
	case DBNotFound:
		return "数据库记录不存在"
	case CamNotExists:
		return "摄像头不存在"
	case DBConflict:
		return "数据库冲突"
	case InternalServerError:
		return "InternalServerError"
	case UserLoginGenTokenFailed:
		return "登录验证失败"
	case UserLoginInvalidUsername:
		return "无效的用户名"
	case DBInternal:
		return "数据库内部错误"
	case EventResReadFailed:
		return "事件资源读取失败"
	case OTALicenseValidate:
		return "证书查询失败"
	case OTAHandleUpload:
		return "处理文件上传失败"
	case OTAFetchRemoteFirmware:
		return "获取远端固件资源错误"
	case OTASubmitFailed:
		return "升级任务提交失败"
	case OTAPadModelInfoFailed:
		return "获取模型信息失败"
	case OTARemoveModelFailed:
		return "移除模型包失败"
	case OTAUpgradeFailed:
		return "升级失败"
	case CamAnalyzeReqError:
		return "分析端请求错误"
	case LicenseValidateByAnalysis:
		return "证书验证失败"
	case LicenseInstallFailed:
		return "证书安装失败"
	case LicenseFileNotExist:
		return "证书文件不存在"
	case LicenseFileHashFailed:
		return "证书文件哈希失败"
	case UploadHandleFailed:
		return "上传文件处理失败"
	}
	return "-"
}

func (v err) CanBeTalkError() bool {
	switch v {
	case BadRequestError:
		return false
	case OTAInvalidParam:
		return true
	case CamElectricFenceInvalid:
		return true
	case CamEnableSectionInvalid:
		return true
	case CamLonLatInvalid:
		return true
	case SysConfigLonLatInvalid:
		return true
	case SysConfigInvalid:
		return true
	case UnauthorizedError:
		return false
	case UserLoginInvalidCaptcha:
		return true
	case UserLoginInvalidToken:
		return true
	case UserLoginInvalidPassword:
		return true
	case ForbiddenError:
		return false
	case OTABusy:
		return true
	case OTADiskLimited:
		return true
	case OTAUnsupportedModel:
		return true
	case OTAUploadLimited:
		return true
	case OTAMd5Checksum:
		return true
	case OTAInvalidVersion:
		return true
	case OTAMaxModelInstall:
		return true
	case OTAInvalidURL:
		return true
	case CamSceneConfLimited:
		return true
	case CamConfLimited:
		return true
	case UploadBusy:
		return true
	case UploadFileSizeLimited:
		return true
	case UploadDiskLimited:
		return true
	case SysConfigPermissionDenied:
		return true
	case NotFoundError:
		return false
	case DBNotFound:
		return true
	case CamNotExists:
		return true
	case DBConflict:
		return true
	case InternalServerError:
		return false
	case UserLoginGenTokenFailed:
		return true
	case UserLoginInvalidUsername:
		return true
	case DBInternal:
		return true
	case EventResReadFailed:
		return true
	case OTALicenseValidate:
		return true
	case OTAHandleUpload:
		return true
	case OTAFetchRemoteFirmware:
		return true
	case OTASubmitFailed:
		return true
	case OTAPadModelInfoFailed:
		return true
	case OTARemoveModelFailed:
		return true
	case OTAUpgradeFailed:
		return true
	case CamAnalyzeReqError:
		return true
	case LicenseValidateByAnalysis:
		return true
	case LicenseInstallFailed:
		return true
	case LicenseFileNotExist:
		return true
	case LicenseFileHashFailed:
		return true
	case UploadHandleFailed:
		return true
	}
	return false
}
