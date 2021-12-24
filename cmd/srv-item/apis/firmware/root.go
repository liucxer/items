package firmware

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var Root = courier.NewRouter(httptransport.Group("/firmware"))

func init() {
	Root.Register(courier.NewRouter(&ListFirmware{}))
	Root.Register(courier.NewRouter(&GetFirmwareInfo{}))
	Root.Register(courier.NewRouter(&GetLatestFirmware{}))
	Root.Register(courier.NewRouter(&UpdateFirmware{}))
	Root.Register(courier.NewRouter(&ReleaseFirmware{}))
	Root.Register(courier.NewRouter(&RevokeReleaseFirmware{}))
	Root.Register(courier.NewRouter(&DeleteFirmware{}))
	Root.Register(courier.NewRouter(&CreateFirmware{}))
}
