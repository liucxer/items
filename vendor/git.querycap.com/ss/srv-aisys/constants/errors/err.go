package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/statuserror"
	"github.com/mattn/go-sqlite3"
)

//go:generate tools gen error err
type err int

func (err) ServiceCode() int {
	return 999 * 1e3
}

func (v err) Err() *Err {
	return &Err{
		err: v,
		se:  v.StatusErr(),
	}
}

type Err struct {
	err
	se *statuserror.StatusErr
}

func (v *Err) StatusErr() *statuserror.StatusErr {
	if v.se == nil {
		v.se = v.err.StatusErr()
	}
	return v.se
}

func (v *Err) WithMsg(msg interface{}) *Err {
	if v.se == nil {
		v.se = v.err.StatusErr()
	}
	v.se.Msg = fmt.Sprint(msg)
	return v
}

func (v *Err) WithDesc(desc interface{}) *Err {
	if v.se == nil {
		v.se = v.err.StatusErr()
	}
	v.se.Desc = fmt.Sprint(desc)
	return v
}

func (v *Err) Msg() string {
	if v.se == nil {
		v.se = v.err.StatusErr()
	}
	return v.se.Msg
}

func (v *Err) Desc() string {
	if v.se == nil {
		v.se = v.err.StatusErr()
	}
	return v.se.Desc
}

var New = errors.New

// 400
const (
	BadRequestError err = http.StatusBadRequest*1e6 + iota + 1
	// @errTalk 升级请求参数非法
	OTAInvalidParam
	// @errTalk 电子围栏参数错误
	CamElectricFenceInvalid
	// @errTalk 场景分析时段数据错误
	CamEnableSectionInvalid
	// @errTalk 相机经纬度数据错误 纬度:[-90,90];经度:[-180,180]
	CamLonLatInvalid
	// @errTalk 设备经纬度数据错误 纬度:[-90,90];经度:[-180,180]
	SysConfigLonLatInvalid
	// @errTalk 配置项不合法
	SysConfigInvalid
)

// 401
const (
	UnauthorizedError err = http.StatusUnauthorized*1e6 + iota + 1
	// @errTalk 验证码错误
	UserLoginInvalidCaptcha
	// @errTalk 登录已过期
	UserLoginInvalidToken
	// @errTalk 密码错误
	UserLoginInvalidPassword
)

// 403
const (
	// Forbidden
	ForbiddenError err = http.StatusForbidden*1e6 + iota + 1
	// @errTalk 升级任务忙
	OTABusy
	// @errTalk 磁盘不足
	OTADiskLimited
	// @errTalk 证书不支持该模型
	OTAUnsupportedModel
	// @errTalk 上传文件过大
	OTAUploadLimited
	// @errTalk md5校验失败
	OTAMd5Checksum
	// @errTalk 低于当前版本
	OTAInvalidVersion
	// @errTalk 可安装达到上限
	OTAMaxModelInstall
	// @errTalk 无效链接
	OTAInvalidURL
	// @errTalk 场景配置达到上限
	CamSceneConfLimited
	// @errTalk 摄像机配置达到上限
	CamConfLimited
	// @errTalk 上传任务繁忙
	UploadBusy
	// @errTalk 上传文件大小限制
	UploadFileSizeLimited
	// @errTalk 上传任务磁盘不足
	UploadDiskLimited
	// @errTalk 无系统配置权限
	SysConfigPermissionDenied
)

// 404
const (
	// NotFound
	NotFoundError err = http.StatusNotFound*1e6 + iota + 1
	// @errTalk 数据库记录不存在
	DBNotFound
	// @errTalk 摄像头不存在
	CamNotExists
)

// 409
const (
	// @errTalk 数据库冲突
	DBConflict err = http.StatusConflict*1e6 + iota + 1
)

// 500
const (
	// InternalServerError
	InternalServerError err = http.StatusInternalServerError*1e6 + iota + 1
	// @errTalk 登录验证失败
	UserLoginGenTokenFailed
	// @errTalk 无效的用户名
	UserLoginInvalidUsername
	// @errTalk 数据库内部错误
	DBInternal
	// @errTalk 事件资源读取失败
	EventResReadFailed
	// @errTalk 证书查询失败
	OTALicenseValidate
	// @errTalk 处理文件上传失败
	OTAHandleUpload
	// @errTalk 获取远端固件资源错误
	OTAFetchRemoteFirmware
	// @errTalk 升级任务提交失败
	OTASubmitFailed
	// @errTalk 获取模型信息失败
	OTAPadModelInfoFailed
	// @errTalk 移除模型包失败
	OTARemoveModelFailed
	// @errTalk 升级失败
	OTAUpgradeFailed
	// @errTalk 分析端请求错误
	CamAnalyzeReqError
	// @errTalk 证书验证失败
	LicenseValidateByAnalysis
	// @errTalk 证书安装失败
	LicenseInstallFailed
	// @errTalk 证书文件不存在
	LicenseFileNotExist
	// @errTalk 证书文件哈希失败
	LicenseFileHashFailed
	// @errTalk 上传文件处理失败
	UploadHandleFailed
)

func DBError(err error) error {
	if err == nil {
		return nil
	}
	if e := sqlx.DBErr(err); e.IsNotFound() {
		return DBNotFound.Err()
	} else if e.IsConflict() {
		return DBConflict.Err()
	} else {
		return DBInternal.Err().WithDesc(err.Error())
	}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return DBError(err).(*Err).err == DBNotFound
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	return DBError(err).(*Err).err == DBConflict
}

func IsSQLiteDBCorrupted(err error) bool {
	if e, ok := err.(sqlite3.Error); ok && e.Code == sqlite3.ErrCorrupt {
		return true
	}
	return false
}
