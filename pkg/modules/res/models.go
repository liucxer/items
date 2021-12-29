package res

import (
	"mime/multipart"

	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

type UploadReq struct {
	File *multipart.FileHeader `name:"file"`
	Info models.ResBase        `name:"info"`
}

type ListReq struct {
	depends.Pager
}

func (r *ListReq) Additions() (additions builder.Additions) {
	rcd := &models.Res{}
	return r.ToAddition(builder.DescOrder(rcd.FieldCreatedAt()))
}

type ListRsp struct {
	Data  []models.Res `json:"data"`
	Total int          `json:"total"`
}
