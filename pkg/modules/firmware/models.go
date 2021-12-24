package firmware

import (
	"mime/multipart"

	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

type CreateReq struct {
	File *multipart.FileHeader `name:"file"`
	Info CreateInfo            `name:"info"`
}

type CreateInfo struct {
	models.FirmwareBase
	models.FirmwareVersion `json:"-"`
	Filename               string       `json:"filename"`            // 固件文件名
	Version                string       `json:"version"`             // 版本号(须符合语意版本号规范x.y.z-...)
	IsRelease              depends.BOOL `json:"isRelease,omitempty"` // 是否立即发布(默认立即发布true)
}

type ListReq struct {
	Name string       `in:"query" name:"name,omitempty"`
	Sort depends.Sort `in:"query" name:"sort,omitempty" default:"releaseAt" validate:"@string{createdAt,releaseAt,major,minor,revision,identifier}{,!asc}"`
	depends.Pager
}

func (r *ListReq) Condition() builder.SqlCondition {
	var (
		cond []builder.SqlCondition
		rcd  = &models.Firmware{}
	)
	if r.Name != "" {
		cond = append(cond, rcd.FieldName().Like(r.Name))
	}
	return builder.And(cond...)
}

func (r *ListReq) Additions() builder.Additions {
	var (
		rcd       = &models.Firmware{}
		additions builder.Additions
	)
	additions = append(additions,
		builder.OrderBy(r.Sort.OrderFor(global.Database(), rcd)))
	if r.Size < 0 {
		additions = append(additions,
			builder.Limit(r.Size).Offset(r.Offset()))
	}
	return additions
}

type ListRsp struct {
	Data  []RspData `json:"data"`
	Total int       `json:"total"`
}

type RspData struct {
	models.Firmware
	Md5     string `json:"md5"`
	Version string `json:"version"`
}

var Controller = &Ctrl{dbe: global.Database()}

type LatestFirmware struct {
	RspData
	URL string `json:"url"`
}
