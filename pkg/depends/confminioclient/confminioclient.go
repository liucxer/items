package confminioclient

import (
	"log"
	"time"

	"github.com/minio/minio-go/v6"
)

type MinioClient struct {
	Endpoint        string        `env:",upstream"`
	AccessKey       string        `env:""`
	SecretKey       string        `env:""`
	UseSSL          bool          `env:""`
	Bucket          string        `env:""`
	Location        string        `env:""`
	ExpiredDuration time.Duration `env:""`
	*minio.Client
}

func (v *MinioClient) SetDefaults() {}

func (v *MinioClient) Init() {
	cli, err := minio.New(v.Endpoint, v.AccessKey, v.SecretKey, v.UseSSL)
	if err != nil {
		log.Panic("创建 MinIO 客户端失败", err)
		return
	}
	log.Printf("创建 MinIO 客户端成功")

	err = cli.MakeBucket(v.Bucket, v.Location)
	if err != nil {
		// 检查存储桶是否已经存在
		exists, err := cli.BucketExists(v.Bucket)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", v.Bucket)
		} else {
			log.Fatalln("查询存储桶状态异常", err)
		}
	}
	log.Printf("创建存储桶 %s 成功", v.Bucket)
	v.Client = cli
}

func (v *MinioClient) Put(object, filename string) error {
	_, err := v.Client.FPutObject(v.Bucket, object, filename, minio.PutObjectOptions{})
	return err
}

func (v *MinioClient) Stat(object string) (minio.ObjectInfo, error) {
	return v.Client.StatObject(v.Bucket, object, minio.StatObjectOptions{})
}

func (v *MinioClient) GetURL(host, object string) (string, error) {
	_, err := v.Client.StatObject(v.Bucket, object, minio.StatObjectOptions{})
	if err != nil {
		return "", err
	}
	u, err := v.Client.PresignedGetObject(v.Bucket, object, v.ExpiredDuration, nil)
	if err != nil {
		return "", err
	}
	u.Host = host
	return u.String(), nil
}

func (v *MinioClient) Delete(object string) error {
	return v.Client.RemoveObject(v.Bucket, object)
}

func init() {
}
