package models

import (
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/enums"
)

//go:generate tools gen model2 Res --database DB --with-comments
// Res 静态资源
// @def primary ID
// @def unique_index UI_res_id ResID
// @def unique_index UI_md5    Md5

type Res struct {
	PrimaryID
	ResRef
	ResBase
	ResExt
	OperationTimes
}

type ResRef struct {
	ResID depends.SFID `db:"f_res_id" json:"resID"` // 资源ID
}

type ResBase struct {
	Type     enums.ResType `db:"f_type,default='0'"    json:"type"`           // 资源类型
	Info     string        `db:"f_info,default=''"     json:"info,omitempty"` // 资源描述信息 用于扩展资源信息 如版本号
	Filename string        `db:"f_filename,default=''" json:"filename"`       // 文件名用于回显资源
}

type ResExt struct {
	URL string `db:"-"     json:"url"` // 资源访问地址
	Md5 string `db:"f_md5" json:"md5"` // 资源MD5
}
