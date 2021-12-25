package items

import (
	"context"

	"github.com/go-courier/httptransport/httpx"
	"github.com/saitofun/items/pkg/models"
	"github.com/saitofun/items/pkg/modules/item"
)

type ListItemsByCode struct {
	httpx.MethodGet `summary:"根据条目代码获取条目及子条目信息" path:"/list/:code"`
	Code            string `in:"path" name:"code"`
}

func (r *ListItemsByCode) Output(ctx context.Context) (interface{}, error) {
	return item.Controller.ListByCode(r.Code)
}

type ListItems struct {
	httpx.MethodGet `summary:"根据条件获取条目列表" path:"/list"`
	item.ListReq
}

func (r *ListItems) Output(ctx context.Context) (interface{}, error) {
	return item.Controller.List(&r.ListReq)
}

type GetItemByCode struct {
	httpx.MethodGet `summary:"根据条目代码获取条目信息" path:"/info/:code"`
	Code            string `in:"path" name:"code"`
}

func (r *GetItemByCode) Output(ctx context.Context) (interface{}, error) {
	return item.Controller.GetByCode(r.Code)
}

type CreateItem struct {
	httpx.MethodPost `summary:"创建条目" path:""`
	models.ItemBase  `in:"body"`
}

func (r *CreateItem) Output(ctx context.Context) (interface{}, error) {
	return item.Controller.CreateItem(&r.ItemBase)
}

type UpdateItemByCode struct {
	httpx.MethodPut `summary:"根据条目代码更新条目" path:"/:code"`
	Code            string `in:"path" name:"code"`
	models.ItemBase `in:"body"`
}

func (r *UpdateItemByCode) Output(ctx context.Context) (interface{}, error) {
	return nil, item.Controller.UpdateItem(r.Code, &r.ItemBase)
}

type DeleteItemByCode struct {
	httpx.MethodDelete `summary:"删除代码或父级代码为Code的所有条目" path:"/:code"`
	Code               string `in:"path" name:"code"`
}

func (r *DeleteItemByCode) Output(ctx context.Context) (interface{}, error) {
	return nil, item.Controller.DeleteByCode(r.Code)
}
