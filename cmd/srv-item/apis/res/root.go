package res

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var Root = courier.NewRouter(httptransport.Group("/res"))

func init() {
	Root.Register(courier.NewRouter(&GetResourceByResID{}))
	Root.Register(courier.NewRouter(&DeleteResourceByResID{}))
	Root.Register(courier.NewRouter(&UploadResource{}))
}
