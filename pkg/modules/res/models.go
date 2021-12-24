package res

import (
	"mime/multipart"

	"github.com/saitofun/items/pkg/models"
)

type UploadReq struct {
	File *multipart.FileHeader `name:"file"`
	Info models.ResBase        `name:"info"`
}
