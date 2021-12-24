package item

import (
	"github.com/go-courier/metax"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
)

type Ctrl struct {
	dbe sqlx.DBExecutor
	metax.Ctx
}

func (c *Ctrl) CreateItem(r *models.ItemBase) (*models.Item, error) {
	rcd := &models.Item{
		ItemRef:  models.ItemRef{ItemID: depends.GenUUID()},
		ItemBase: *r,
	}
	err := rcd.Create(c.dbe)
	if err != nil {
		return nil, err
	}
	err = rcd.FetchByCode(c.dbe)
	if err != nil {
		return nil, err
	}

	return rcd, nil
}

func (c *Ctrl) UpdateItem(code string, r *models.ItemBase) error {
	old := &models.Item{ItemBase: models.ItemBase{Code: code}}
	if err := old.FetchByCode(c.dbe); err != nil {
		return err
	}
	rcd := &models.Item{ItemRef: models.ItemRef{ItemID: old.ItemID}, ItemBase: *r}
	return rcd.UpdateByItemIDWithStruct(c.dbe)
}

func (c *Ctrl) ListByCode(code string) (*ListRsp, error) {
	var (
		err error
		rcd = &models.Item{ItemBase: models.ItemBase{Code: code}}
		rsp = &ListRsp{}
	)
	rsp.Data, err = rcd.List(c.dbe, builder.And(
		rcd.FieldCode().Eq(code),
		rcd.FieldParentCode().Eq(code),
	), nil)
	if err != nil {
		return nil, err
	}
	rsp.Total = len(rsp.Data)
	return rsp, nil
}

func (c *Ctrl) List(r *ListReq) (*ListRsp, error) {
	var (
		err error
		rcd = &models.Item{}
		rsp = &ListRsp{}
	)
	rsp.Data, err = rcd.List(c.dbe, r.Condition(), r.Additions()...)
	if err != nil {
		return nil, err
	}
	rsp.Total, err = rcd.Count(c.dbe, r.Condition())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Ctrl) GetByCode(code string) (*models.Item, error) {
	rcd := &models.Item{ItemBase: models.ItemBase{Code: code}}
	err := rcd.FetchByCode(c.dbe)
	if err != nil {
		return nil, err
	}
	return rcd, nil
}

func (c *Ctrl) DeleteByCode(code string) error {
	rcd := &models.Item{}
	tab := c.dbe.T(rcd)
	_, err := c.dbe.ExecExpr(builder.Delete().From(
		tab,
		builder.Where(
			builder.Or(
				rcd.FieldCode().Eq(code),
				rcd.FieldParentCode().Eq(code),
			),
		),
	))
	return err
}

func (c *Ctrl) UploadAttachment() {}

func (c *Ctrl) UpdateItemIcon() {}

var Controller = &Ctrl{
	dbe: global.Database(),
}
