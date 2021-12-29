package firmware

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var Root = courier.NewRouter(httptransport.Group("/firmware"))
var RootAuth = courier.NewRouter(httptransport.Group("/firmware"))

func init() {
	Root.Register(courier.NewRouter(&ListFirmware{}))
	Root.Register(courier.NewRouter(&GetFirmwareInfo{}))
	Root.Register(courier.NewRouter(&GetLatestFirmware{}))
	RootAuth.Register(courier.NewRouter(&UpdateFirmware{}))
	RootAuth.Register(courier.NewRouter(&ReleaseFirmware{}))
	RootAuth.Register(courier.NewRouter(&RevokeReleaseFirmware{}))
	RootAuth.Register(courier.NewRouter(&DeleteFirmware{}))
	RootAuth.Register(courier.NewRouter(&CreateFirmware{}))
}
