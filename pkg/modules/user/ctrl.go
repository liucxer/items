package user

import (
	"time"

	"github.com/go-courier/metax"
	"github.com/go-courier/sqlx/v2"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/errors"
	"github.com/saitofun/items/pkg/models"
)

type Ctrl struct {
	dbe sqlx.DBExecutor
	metax.Ctx
}

func (c *Ctrl) UpdatePassword(username string, r *UpdatePasswordReq) error {
	var (
		rcd = &models.User{UserBase: models.UserBase{Username: username}}
		err = rcd.FetchByUsername(c.dbe)
	)
	if err != nil {
		return err
	}
	if r.OldPassword != rcd.Password {
		return errors.UnmatchedPassword
	}
	rcd.Password = r.NewPassword
	return rcd.UpdateByUsernameWithStruct(c.dbe)
}

func (c *Ctrl) Login(r *LoginReq) (*LoginRsp, error) {
	var (
		rcd   = &models.User{UserBase: models.UserBase{Username: r.Username}}
		err   = rcd.FetchByUsername(c.dbe)
		token string
		exp   = time.Now().Add(time.Hour).Unix()
	)
	if err != nil {
		return nil, err
	}
	if r.Password != rcd.Password {
		return nil, errors.UnmatchedPassword
	}
	token, err = GenerateToken(r.Username, r.Password)
	if err != nil {
		return nil, errors.GenerateTokenError
	}

	return &LoginRsp{
		Username: r.Username,
		Token:    token,
		Expire:   exp,
	}, nil
}

func (c *Ctrl) CreateUser(r *models.UserBase) (*models.User, error) {
	rcd := &models.User{
		UserRef:  models.UserRef{UserID: depends.GenUUID()},
		UserBase: *r,
	}
	err := rcd.Create(c.dbe)
	if err != nil {
		return nil, err
	}
	err = rcd.FetchByUsername(c.dbe)
	if err != nil {
		return nil, err
	}
	rcd.Password = ""
	return rcd, nil
}

func (c *Ctrl) List(r *ListReq) (*ListRsp, error) {
	var (
		err error
		rcd = &models.User{}
		rsp = &ListRsp{}
	)
	rsp.Data, err = rcd.List(c.dbe, nil, r.Additions()...)
	if err != nil {
		return nil, err
	}
	rsp.Total, err = rcd.Count(c.dbe, nil)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Ctrl) DeleteByUsername(username string) error {
	rcd := &models.User{UserBase: models.UserBase{Username: username}}
	return rcd.DeleteByUsername(c.dbe)
}

func (c *Ctrl) DeleteByID(id depends.SFID) error {
	rcd := &models.User{UserRef: models.UserRef{UserID: id}}
	return rcd.DeleteByUserID(c.dbe)
}

var Controller = &Ctrl{
	dbe: global.Database(),
}
