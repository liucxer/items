package user

import (
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRsp struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expire   int64  `json:"expire"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

type ListReq struct {
	Sort depends.Sort `in:"query" name:"sort,omitempty" default:"createdAt" validate:"@string{createdAt,updatedAt}{,!asc}"`
	depends.Pager
}

func (r *ListReq) Additions() builder.Additions {
	var (
		rcd       = &models.User{}
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
	Data  []models.User `json:"data"`
	Total int           `json:"total"`
}
