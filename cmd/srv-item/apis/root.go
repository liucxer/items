package apis

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
	"github.com/go-courier/httptransport/openapi"
	"github.com/saitofun/items/cmd/srv-item/apis/firmware"
	"github.com/saitofun/items/cmd/srv-item/apis/items"
	"github.com/saitofun/items/cmd/srv-item/apis/middleware"
	"github.com/saitofun/items/cmd/srv-item/apis/res"
	"github.com/saitofun/items/cmd/srv-item/apis/users"
)

var (
	Root = courier.NewRouter(httptransport.Group("/item_manager"))
	v0   = courier.NewRouter(httptransport.Group("/v0"))
	auth = courier.NewRouter(&middleware.Auth{})
)

func init() {
	Root.Register(v0)
	Root.Register(openapi.OpenAPIRouter)
	v0.Register(courier.NewRouter(&users.Login{}))
	v0.Register(auth)
	v0.Register(items.Root)
	auth.Register(users.Root)
	auth.Register(items.RootAuth)
	auth.Register(res.Root)
	auth.Register(firmware.Root)
}
