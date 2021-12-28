package models

import "github.com/saitofun/items/pkg/depends"

//go:generate tools gen model2 Item --database DB --with-comments
// Item 条目信息
// @def primary ID
// @def unique_index UI_item_id          ItemID
// @def unique_index UI_code             Code
// @def unique_index UI_code_parent_code Code ParentCode

type Item struct {
	PrimaryID
	ItemRef
	ItemBase
	OperationTimes
}

type ItemRef struct {
	ItemID depends.SFID `db:"f_item_id" json:"itemID"` // 条目编号(UUID)
}

type ItemBase struct {
	Code        string       `db:"f_code"                      json:"code"`                  // 条目代码
	ParentCode  string       `db:"f_parent_code,default=''"    json:"parentCode,omitempty"`  // 条目上级代码
	Name        string       `db:"f_name"                      json:"name"`                  // 条目名称
	AlphabetZH  string       `db:"f_alphabet_zh,default=''"    json:"alphabetZH,omitempty"`  // 条目中文拼音
	AlphabetEN  string       `db:"f_alphabet_en,default=''"    json:"alphabetEN,omitempty"`  // 条目英文
	ImageResID  depends.SFID `db:"f_image_res_id,default='0'"  json:"imageResID,omitempty"`  // 条目icon资源ID
	RichText    string       `db:"f_rich_text,default=''"      json:"richText,omitempty"`    // 条目富文本
	Link        string       `db:"f_link,default=''"           json:"link,omitempty"`        // 条目跳转链接
	AttachResID depends.SFID `db:"f_attach_res_id,default='0'" json:"attachResID,omitempty"` // 条目附件资源ID
	HasSub      bool         `db:"f_has_sub,default='false'"   json:"hasSub,omitempty"`      // 是否有下一级条目
}
