package users

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var Root = courier.NewRouter(httptransport.Group("/user"))

func init() {
	Root.Register(courier.NewRouter(&Logout{}))
	Root.Register(courier.NewRouter(&CreateUser{}))
	Root.Register(courier.NewRouter(&UpdatePassword{}))
	Root.Register(courier.NewRouter(&DeleteByUsername{}))
	Root.Register(courier.NewRouter(&ListUsers{}))
}
