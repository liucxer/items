package res

import (
	"context"

	types "git.querycap.com/tools/datatypes"
	"github.com/go-courier/httptransport/httpx"
	"github.com/saitofun/items/pkg/modules/res"
)

type UploadResource struct {
	httpx.MethodPost `summary:"上传资源" path:""`
	res.UploadReq    `in:"body" mime:"multipart"`
}

func (r *UploadResource) Output(ctx context.Context) (interface{}, error) {
	return res.Controller.Upload(&r.UploadReq)
}

type GetResourceByResID struct {
	httpx.MethodGet `summary:"根据资源ID获取资源信息" path:"/:resID"`
	ResID           types.SFID `in:"path" name:"resID"`
}

func (r *GetResourceByResID) Output(ctx context.Context) (interface{}, error) {
	return res.Controller.GetByID(r.ResID)
}

type DeleteResourceByResID struct {
	httpx.MethodDelete `summary:"根据资源ID删除资源" path:"/:resID"`
	ResID              types.SFID `in:"path" name:"resID"`
}

func (r *DeleteResourceByResID) Output(ctx context.Context) (interface{}, error) {
	return nil, res.Controller.DeleteByID(r.ResID)
}
