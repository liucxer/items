package users

import (
	"context"

	"github.com/go-courier/httptransport/httpx"
	"github.com/saitofun/items/cmd/srv-item/apis/middleware"
	"github.com/saitofun/items/pkg/models"
	"github.com/saitofun/items/pkg/modules/user"
)

type Login struct {
	httpx.MethodPost `summary:"用户登陆" path:"/login"`
	user.LoginReq    `in:"body"`
}

func (r *Login) Output(ctx context.Context) (interface{}, error) {
	return user.Controller.Login(&r.LoginReq)
}

type Logout struct {
	httpx.MethodPost `summary:"用户登出" path:"/logout"`
}

func (r *Logout) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type CreateUser struct {
	httpx.MethodPost `summary:"创建用户" path:""`
	models.UserBase  `in:"body"`
}

func (r *CreateUser) Output(ctx context.Context) (interface{}, error) {
	return user.Controller.CreateUser(&r.UserBase)
}

type UpdatePassword struct {
	httpx.MethodPut        `summary:"更新当前用户密码" path:"/modify_password"`
	user.UpdatePasswordReq `in:"body"`
}

func (r *UpdatePassword) Output(ctx context.Context) (interface{}, error) {
	return nil, user.Controller.UpdatePassword(
		middleware.GetContext(ctx).Username,
		&r.UpdatePasswordReq,
	)
}

type ListUsers struct {
	httpx.MethodGet `summary:"用户列表" path:""`
	user.ListReq
}

func (r *ListUsers) Output(ctx context.Context) (interface{}, error) {
	return user.Controller.List(&r.ListReq)
}

type DeleteByUsername struct {
	httpx.MethodDelete `summary:"根据用户名删除用户信息" path:"/:username"`
	Username           string `in:"path" name:"username"`
}

func (r *DeleteByUsername) Output(ctx context.Context) (interface{}, error) {
	return nil, user.Controller.DeleteByUsername(r.Username)
}
