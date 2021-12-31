package item

import (
	"git.querycap.com/ss/srv-aisys/constants/errors"
	"github.com/go-courier/metax"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
	"github.com/saitofun/items/pkg/modules/res"
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
		return nil, errors.DBError(err)
	}
	err = rcd.FetchByCode(c.dbe)
	if err != nil {
		return nil, errors.DBError(err)
	}
	return rcd, nil
}

func (c *Ctrl) UpdateItem(code string, r *models.ItemBase) error {
	old := &models.Item{ItemBase: models.ItemBase{Code: code}}
	if err := old.FetchByCode(c.dbe); err != nil {
		return errors.DBError(err)
	}
	tab := c.dbe.T(old)
	expr := builder.Update(tab).Where(old.FieldCode().Eq(code)).Set(
		tab.AssignmentsByFieldValues(builder.FieldValues{
			old.FieldKeyParentCode():  r.ParentCode,
			old.FieldKeyName():        r.Name,
			old.FieldKeyAlphabetZH():  r.AlphabetZH,
			old.FieldKeyAlphabetEN():  r.AlphabetEN,
			old.FieldKeyImageResID():  r.ImageResID,
			old.FieldKeyRichText():    r.RichText,
			old.FieldKeyLink():        r.Link,
			old.FieldKeyAttachResID(): r.AttachResID,
			old.FieldKeyHasSub():      r.HasSub,
		})...)
	_, err := c.dbe.ExecExpr(expr)
	return errors.DBError(err)
}

func (c *Ctrl) ListByCode(code string) (*ListRsp, error) {
	var (
		err error
		rcd = &models.Item{ItemBase: models.ItemBase{Code: code}}
		rsp = &ListRsp{}
		lst []models.Item
	)
	lst, err = rcd.List(c.dbe, builder.And(
		rcd.FieldCode().Eq(code),
		rcd.FieldParentCode().Eq(code),
	), nil)
	if err != nil {
		return nil, errors.DBError(err)
	}

	for i := range lst {
		data := ListData{Item: lst[i]}
		if lst[i].AttachResID != 0 {
			attach, err := res.Controller.GetByID(lst[i].AttachResID)
			if err != nil {
				return nil, err
			}
			data.Attach = &ResData{
				ResBase: attach.ResBase,
				ResExt:  attach.ResExt,
			}
		}
		if lst[i].ImageResID != 0 {
			image, err := res.Controller.GetByID(lst[i].ImageResID)
			if err != nil {
				return nil, err
			}
			data.Image = &ResData{
				ResBase: image.ResBase,
				ResExt:  image.ResExt,
			}
		}
		rsp.Data = append(rsp.Data, data)
	}

	rsp.Total = len(rsp.Data)
	return rsp, nil
}

func (c *Ctrl) List(r *ListReq) (*ListRsp, error) {
	var (
		err error
		rcd = &models.Item{}
		rsp = &ListRsp{}
		lst []models.Item
	)
	lst, err = rcd.List(c.dbe, r.Condition(), r.Additions()...)
	if err != nil {
		return nil, errors.DBError(err)
	}
	rsp.Total, err = rcd.Count(c.dbe, r.Condition())
	if err != nil {
		return nil, errors.DBError(err)
	}
	for i := range lst {
		data := ListData{Item: lst[i]}
		if lst[i].AttachResID != 0 {
			attach, err := res.Controller.GetByID(lst[i].AttachResID)
			if err != nil {
				return nil, err
			}
			data.Attach = &ResData{
				ResBase: attach.ResBase,
				ResExt:  attach.ResExt,
			}
		}
		if lst[i].ImageResID != 0 {
			image, err := res.Controller.GetByID(lst[i].ImageResID)
			if err != nil {
				return nil, err
			}
			data.Image = &ResData{
				ResBase: image.ResBase,
				ResExt:  image.ResExt,
			}
		}
		rsp.Data = append(rsp.Data, data)
	}
	return rsp, nil
}

func (c *Ctrl) GetByCode(code string) (*models.Item, error) {
	rcd := &models.Item{ItemBase: models.ItemBase{Code: code}}
	err := rcd.FetchByCode(c.dbe)
	if err != nil {
		return nil, errors.DBError(err)
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
	return errors.DBError(err)
}

var Controller = &Ctrl{dbe: global.Database()}
