package res

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/go-courier/metax"
	"github.com/go-courier/sqlx/v2"
	"github.com/google/uuid"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends"
	"github.com/saitofun/items/pkg/errors"
	"github.com/saitofun/items/pkg/models"
	"github.com/shirou/gopsutil/disk"
)

type Ctrl struct {
	dbe sqlx.DBExecutor
	metax.Ctx
}

var reserve = int64(100 * 1024 * 1024)

func Upload(file *multipart.FileHeader, dst string, limit int64) (filename string, err error) {
	var (
		fr       io.ReadSeekCloser
		fw       io.WriteCloser
		filesize = int64(0)
	)

	filename = path.Join(global.ResPath, uuid.New().String()+"-"+dst)

	if !IsPathExists(dst) {
		if err = os.Mkdir(dst, 0777); err != nil {
			return
		}
	}

	if fr, err = file.Open(); err != nil {
		return
	}
	defer fr.Close()

	if filesize, err = fr.Seek(0, io.SeekEnd); err != nil {
		return
	}
	if filesize > limit {
		err = errors.New("filesize over limit")
		return
	}
	if stat, _ := disk.Usage(dst); stat == nil || stat.Free < uint64(filesize+reserve) {
		err = errors.New("disk limited")
		return
	}
	_, err = fr.Seek(0, io.SeekStart)
	if err != nil {
		return
	}
	if fw, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666); err != nil {
		return
	}
	defer fw.Close()
	if _, err = io.Copy(fw, fr); err != nil {
		return
	}
	return filename, nil
}

func Md5Hash(path string) ([]byte, error) {
	hash := md5.New()
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err = io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func Md5HashString(path string) (string, error) {
	hash, err := Md5Hash(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}

func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (c *Ctrl) Upload(r *UploadReq) (*models.Res, error) {
	filename, err := Upload(r.File, global.ResPath, 100*1024*1024)
	if err != nil {
		return nil, errors.InternalServerError.WithDes(err)
	}
	defer os.RemoveAll(filename)
	md5, err := Md5HashString(filename)
	if err != nil {
		return nil, errors.InternalServerError.WithDes(err)
	}

	rcd := &models.Res{
		ResBase: models.ResBase{
			Type:     r.Info.Type,
			Info:     r.Info.Info,
			Filename: r.Info.Filename},
		ResExt: models.ResExt{Md5: md5},
	}

	resID := depends.SFID(0)
	if err = rcd.FetchByMd5(c.dbe); err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			resID = depends.GenUUID()
			if err = global.MinioClient.Put(resID.String(), filename); err != nil {
				return nil, errors.UploadStorage.WithDes(err)
			}
			rcd.ResID = resID
			if err = rcd.Create(c.dbe); err != nil {
				return nil, errors.DBError(err)
			}
		} else {
			return nil, errors.DBError(err)
		}
	} else {
		resID = rcd.ResID
	}
	url, err := global.MinioClient.GetURL(global.MinioHost, resID.String())
	if err != nil {
		return nil, errors.GetDownloadLink.WithDes(err)
	}
	rcd.URL = url
	return rcd, nil
}

func (c *Ctrl) GetByID(id depends.SFID) (*models.Res, error) {
	rcd := &models.Res{ResRef: models.ResRef{ResID: id}}
	err := rcd.FetchByResID(c.dbe)
	if err != nil {
		return nil, errors.DBError(err)
	}
	url, err := global.MinioClient.GetURL(global.MinioHost, id.String())
	if err != nil {
		return nil, errors.GetDownloadLink.WithDes(err)
	}
	rcd.URL = url
	return rcd, nil
}

func (c *Ctrl) DeleteByID(id depends.SFID) error {
	rcd := &models.Res{ResRef: models.ResRef{ResID: id}}
	if err := global.MinioClient.Delete(id.String()); err != nil {
		return errors.GetDownloadLink.WithDes(err)
	}
	return errors.DBError(rcd.DeleteByResID(c.dbe))
}

func (c *Ctrl) List(r *ListReq) (ret *ListRsp, err error) {
	rcd := &models.Res{}
	ret = &ListRsp{}
	ret.Data, err = rcd.List(c.dbe, nil, r.Additions()...)
	if err != nil {
		return nil, errors.DBError(err)
	}
	ret.Total, err = rcd.Count(c.dbe, nil)
	if err != nil {
		return nil, errors.DBError(err)
	}
	return
}

var Controller = &Ctrl{
	dbe: global.Database(),
}
