package items

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var Root = courier.NewRouter(httptransport.Group("/item"))
var RootAuth = courier.NewRouter(httptransport.Group("/item"))

func init() {
	Root.Register(courier.NewRouter(&ListItems{}))
	Root.Register(courier.NewRouter(&ListItemsByCode{}))
	Root.Register(courier.NewRouter(&GetItemByCode{}))
	RootAuth.Register(courier.NewRouter(&CreateItem{}))
	RootAuth.Register(courier.NewRouter(&UpdateItemByCode{}))
	RootAuth.Register(courier.NewRouter(&DeleteItemByCode{}))
}
