package item

import (
	"git.querycap.com/ss/srv-aisys/depends/pager"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

type ListReq struct {
	Codes       []string     `in:"query" name:"code,omitempty"`
	ParentCodes []string     `in:"query" name:"parentCode,omitempty"`
	Sort        depends.Sort `in:"query" name:"sort,omitempty" default:"createdAt" validate:"@string{createdAt,updatedAt}{,!asc}"`
	pager.Pager
}

func (v *ListReq) Condition() builder.SqlCondition {
	var (
		rcd        = &models.Item{}
		conditions []builder.SqlCondition
	)
	if len(v.Codes) > 0 {
		conditions = append(conditions, rcd.FieldCode().In(v.Codes))
	}
	if len(v.ParentCodes) > 0 {
		conditions = append(conditions, rcd.FieldParentCode().In(v.ParentCodes))
	}
	return builder.And(conditions...)
}

func (v *ListReq) Additions() builder.Additions {
	rcd := &models.Item{}
	additions := make(builder.Additions, 0)
	additions = append(additions,
		builder.OrderBy(v.Sort.OrderFor(global.Database(), rcd)))
	if v.Size < 0 {
		additions = append(additions,
			builder.Limit(v.Size).Offset(v.Offset()))
	}
	return additions
}

type ListRsp struct {
	Data  []ListData `json:"data"`
	Total int        `json:"total"`
}

type ResData struct {
	models.ResBase
	models.ResExt
}

type ListData struct {
	models.Item
	Image  *ResData `json:"image,omitempty"`
	Attach *ResData `json:"attach,omitempty"`
}
