package firmware

import (
	"context"

	"github.com/go-courier/httptransport/httpx"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/models"
	"github.com/saitofun/items/pkg/modules/firmware"
)

type ListFirmware struct {
	httpx.MethodGet `summary:"固件列表" path:"/list"`
	firmware.ListReq
}

func (r *ListFirmware) Output(ctx context.Context) (interface{}, error) {
	return firmware.Controller.ListFirmware(&r.ListReq)
}

type GetFirmwareInfo struct {
	httpx.MethodGet `summary:"获取固件信息" path:"/info/:firmwareID"`
	FirmwareID      depends.SFID `in:"path" name:"firmwareID"`
}

func (r *GetFirmwareInfo) Output(ctx context.Context) (interface{}, error) {
	return firmware.Controller.GetFirmware(r.FirmwareID)
}

type GetLatestFirmware struct {
	httpx.MethodGet `summary:"获取最新固件" path:"/latest"`
	Version         string `in:"query" name:"version,omitempty"`
}

func (r *GetLatestFirmware) Output(ctx context.Context) (interface{}, error) {
	version, err := models.ParseFirmwareVersion(r.Version)
	if err != nil {
		return nil, err
	}
	return firmware.Controller.GetLatest(version)
}

type UpdateFirmware struct {
	httpx.MethodPut     `summary:"更新固件信息" path:"/info/:firmwareID"`
	FirmwareID          depends.SFID `in:"path" name:"firmwareID"`
	firmware.CreateInfo `in:"body"`
}

func (r *UpdateFirmware) Output(ctx context.Context) (interface{}, error) {
	return nil, firmware.Controller.UpdateFirmware(r.FirmwareID, &r.CreateInfo)
}

type ReleaseFirmware struct {
	httpx.MethodPut `summary:"发布固件" path:"/release/:firmwareID"`
	FirmwareID      depends.SFID `in:"path" name:"firmwareID"`
}

func (r *ReleaseFirmware) Output(ctx context.Context) (interface{}, error) {
	return nil, firmware.Controller.Release(r.FirmwareID)
}

type RevokeReleaseFirmware struct {
	httpx.MethodPut `summary:"撤销发布固件" path:"/revoke_release/:firmwareID"`
	FirmwareID      depends.SFID `in:"path" name:"firmwareID"`
}

func (r *RevokeReleaseFirmware) Output(ctx context.Context) (interface{}, error) {
	return nil, firmware.Controller.RevokeRelease(r.FirmwareID)
}

type DeleteFirmware struct {
	httpx.MethodDelete `summary:"删除固件" path:"/:firmwareID"`
	FirmwareID         depends.SFID `in:"path" name:"firmwareID"`
}

func (r *DeleteFirmware) Output(ctx context.Context) (interface{}, error) {
	return nil, firmware.Controller.Delete(r.FirmwareID)
}

type CreateFirmware struct {
	httpx.MethodPost   `summary:"新增固件" path:""`
	firmware.CreateReq `in:"body" mime:"multipart"`
}

func (r *CreateFirmware) Output(ctx context.Context) (interface{}, error) {
	version, err := models.ParseFirmwareVersion(r.Info.Version)
	if err != nil {
		return nil, err
	}
	r.CreateReq.Info.FirmwareVersion = *version
	return firmware.Controller.Create(&r.CreateReq)
}
