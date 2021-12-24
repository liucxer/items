package firmware

import (
	"encoding/json"
	"time"

	"github.com/go-courier/metax"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/enums"
	"github.com/saitofun/items/pkg/models"
	"github.com/saitofun/items/pkg/modules/res"
)

type Ctrl struct {
	dbe sqlx.DBExecutor
	metax.Ctx
}

func (c *Ctrl) Create(r *CreateReq) (*RspData, error) {
	var (
		rcd     *models.Res
		info, _ = json.Marshal(r.Info)
		pkg     *models.Firmware
	)

	err := sqlx.NewTasks(c.dbe).With(
		func(db sqlx.DBExecutor) error {
			var e error
			rcd, e = res.Controller.Upload(&res.UploadReq{
				File: r.File,
				Info: models.ResBase{
					Type:     enums.RES_TYPE__APK,
					Info:     string(info),
					Filename: r.Info.Filename,
				},
			})
			return e
		},
		func(db sqlx.DBExecutor) error {
			pkg = &models.Firmware{
				FirmwareRef:     models.FirmwareRef{FirmwareID: depends.GenUUID()},
				ResRef:          models.ResRef{ResID: rcd.ResID},
				FirmwareVersion: r.Info.FirmwareVersion,
				FirmwareBase:    r.Info.FirmwareBase,
			}
			if r.Info.IsRelease == depends.BOOL(0) ||
				r.Info.IsRelease == depends.T {
				pkg.ReleaseAt = depends.Timestamp(time.Now())
			}
			return pkg.Create(c.dbe)
		},
	).Do()
	if err != nil {
		return nil, err
	}
	return &RspData{
		Firmware: *pkg,
		Md5:      rcd.Md5,
		Version:  pkg.FirmwareVersion.String(),
	}, nil
}

func (c *Ctrl) UpdateFirmware(id depends.SFID, info *CreateInfo) error {
	pkg := &models.Firmware{
		FirmwareRef:     models.FirmwareRef{FirmwareID: id},
		FirmwareVersion: info.FirmwareVersion,
		FirmwareBase:    info.FirmwareBase,
	}
	if info.IsRelease == depends.T || info.IsRelease == depends.BOOL(0) {
		pkg.ReleaseAt = depends.Timestamp(time.Now())
		return pkg.UpdateByFirmwareIDWithStruct(c.dbe,
			pkg.FieldKeyResID())
	}
	return pkg.UpdateByFirmwareIDWithStruct(c.dbe,
		pkg.FieldKeyResID(), pkg.FieldKeyReleaseAt())
}

func (c *Ctrl) Release(id depends.SFID) error {
	pkg := &models.Firmware{}
	tab := c.dbe.T(pkg)
	_, err := c.dbe.ExecExpr(builder.
		Update(tab).
		Where(pkg.FieldFirmwareID().Eq(id)).
		Set(tab.AssignmentsByFieldValues(builder.FieldValues{
			pkg.FieldKeyReleaseAt(): depends.Timestamp(time.Now()),
			pkg.FieldKeyUpdatedAt(): depends.Timestamp(time.Now()),
		})...))
	return err
}

func (c *Ctrl) RevokeRelease(id depends.SFID) error {
	pkg := &models.Firmware{}
	tab := c.dbe.T(pkg)
	_, err := c.dbe.ExecExpr(builder.
		Update(tab).
		Where(pkg.FieldFirmwareID().Eq(id)).
		Set(tab.AssignmentsByFieldValues(builder.FieldValues{
			pkg.FieldKeyReleaseAt(): 0,
			pkg.FieldKeyUpdatedAt(): depends.Timestamp(time.Now()),
		})...))
	return err
}

func (c *Ctrl) ListFirmware(r *ListReq) (*ListRsp, error) {
	var (
		err error
		ret = &ListRsp{}
		rcd = &models.Firmware{}
		lst []models.Firmware
	)

	lst, err = rcd.List(c.dbe, r.Condition(), r.Additions()...)
	if err != nil {
		return nil, err
	}
	ret.Total, err = rcd.Count(c.dbe, r.Condition())
	if err != nil {
		return nil, err
	}
	for _, v := range lst {
		resRef := &models.Res{}
		resRef.ResID = v.ResID
		if err = resRef.FetchByResID(c.dbe); err != nil {
			return nil, err
		}
		ret.Data = append(ret.Data, RspData{
			Firmware: v,
			Md5:      resRef.Md5,
			Version:  v.FirmwareVersion.String(),
		})
	}
	return ret, err
}

func (c *Ctrl) GetFirmware(id depends.SFID) (*models.Firmware, error) {
	pkg := &models.Firmware{
		FirmwareRef: models.FirmwareRef{FirmwareID: id},
	}
	err := pkg.FetchByFirmwareID(c.dbe)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func (c *Ctrl) Delete(id depends.SFID) error {
	return (&models.Firmware{
		FirmwareRef: models.FirmwareRef{FirmwareID: id},
	}).DeleteByFirmwareID(c.dbe)
}

func (c *Ctrl) GetLatest(cur *models.FirmwareVersion) (*LatestFirmware, error) {
	var (
		pkg = &models.Firmware{}
		ret []models.Firmware
	)

	err := c.dbe.QueryExprAndScan(
		builder.Select(nil).From(
			c.dbe.T(pkg),
			builder.Where(pkg.FieldReleaseAt().Gt(0)),
			builder.OrderBy(
				builder.DescOrder(pkg.FieldMajor()),
				builder.DescOrder(pkg.FieldMinor()),
				builder.DescOrder(pkg.FieldPatch()),
			),
			builder.Limit(1),
		),
		&ret,
	)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, nil
	}

	v := ret[0]
	if v.Major < cur.Major {
		return nil, nil
	}
	if v.Minor < cur.Minor {
		return nil, nil
	}
	if v.Patch < cur.Patch {
		return nil, nil
	}
	if v.Major == cur.Major && v.Minor == v.Minor && v.Patch == v.Patch {
		return nil, nil
	}

	resource, err := res.Controller.GetByID(ret[0].ResID)
	if err != nil {
		return nil, err
	}

	return &LatestFirmware{
		RspData: RspData{
			Firmware: ret[0],
			Md5:      resource.Md5,
			Version:  ret[0].FirmwareVersion.String(),
		},
		URL: resource.URL,
	}, nil
}
