package models

import (
	"fmt"

	"github.com/go-courier/semver"
	"github.com/saitofun/items/pkg/depends"
)

//go:generate tools gen model2 Firmware --database DB --with-comments
// @def primary ID
// @def unique_index UI_firmware_id FirmwareID
// @def unique_index UI_res_id      ResID
// @def unique_index UI_version     Name Major Minor Patch Identifier ResID

type Firmware struct {
	PrimaryID
	FirmwareRef
	ResRef
	FirmwareVersion
	FirmwareBase
	FirmwareRelease
	OperationTimes
}

type FirmwareRef struct {
	FirmwareID depends.SFID `db:"f_firmware_id" json:"firmwareID"` // 固件ID
}

type FirmwareVersion struct {
	Major      uint64 `db:"f_major,default='0'"     json:"major,omitempty"`      // 主版本号
	Minor      uint64 `db:"f_minor,default='0'"     json:"minor,omitempty"`      // 次版本号
	Patch      uint64 `db:"f_patch,default='0'"     json:"patch,omitempty"`      // 修订号
	Identifier string `db:"f_identifier,default=''" json:"identifier,omitempty"` // 修饰符
}

func (v FirmwareVersion) String() string {
	ret := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Identifier != "" {
		ret += "-" + v.Identifier
	}
	return ret
}

func ParseFirmwareVersion(s string) (*FirmwareVersion, error) {
	ret := &FirmwareVersion{}
	ver, err := semver.ParseVersion(s)
	if err != nil {
		return nil, err
	}
	ret.Major = ver.Major()
	ret.Minor = ver.Minor()
	ret.Patch = ver.Patch()
	if prerelease := ver.Prerelease(); prerelease != "" {
		ret.Identifier += "-" + prerelease
	}
	if meta := ver.Metadata(); meta != "" {
		ret.Identifier += "+" + meta
	}
	return ret, nil
}

type FirmwareBase struct {
	Name string `db:"f_name"            json:"name"`           // 固件名称
	Desc string `db:"f_desc,default=''" json:"desc,omitempty"` // 固件描述
}

type FirmwareRelease struct {
	ReleaseAt depends.Timestamp `db:"f_release_at,default='0'" json:"releaseAt,omitempty"` // 发布时间
}
