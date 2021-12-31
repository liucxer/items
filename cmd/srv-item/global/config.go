package global

import (
	"net"
	"os"

	"git.querycap.com/tools/confpostgres"
	"git.querycap.com/tools/scaffold/pkg/appinfo"
	"git.querycap.com/tools/svcutil/confhttp"
	"git.querycap.com/tools/svcutil/conflogger"
	"github.com/go-courier/courier"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/migration"
	"github.com/saitofun/items/pkg/depends/confminioclient"
	"github.com/saitofun/items/pkg/models"
)

var (
	server      = &confhttp.Server{}
	database    = &confpostgres.Postgres{Database: models.DB}
	ResPath     string
	MinioClient confminioclient.MinioClient

	App          *appinfo.AppCtx
	MinioHost    string
	MinioFwdPort string
	iface        string
)

func init() {
	var config = &struct {
		Log          *conflogger.Log
		Server       *confhttp.Server
		DB           *confpostgres.Postgres
		MinioClient  *confminioclient.MinioClient
		ResPath      *string `env:""`
		MinioFwdPort *string `env:""`
		DefaultIface *string `env:""`
	}{
		Server:       server,
		DB:           database,
		ResPath:      &ResPath,
		MinioClient:  &MinioClient,
		MinioFwdPort: &MinioFwdPort,
		DefaultIface: &iface,
	}

	// pwd, _ := os.Getwd()

	App = appinfo.New(
		appinfo.WithName("srv-item"),
		// appinfo.WithMainRoot("."),
	)
	App.ConfP(config)

	os.MkdirAll(*config.ResPath, 0777)

	confhttp.RegisterCheckerFromStruct(config)
}

func init() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range interfaces {
		if i.Name == iface {
			addrs, _ := i.Addrs()
			if len(addrs) == 0 {
				panic("addrs[0]")
			}
			if va, ok := addrs[0].(*net.IPNet); ok {
				MinioHost = va.IP.String() + ":" + MinioFwdPort
			} else {
				panic("not IP net")
			}
		}
	}
	if MinioHost == "" {
		panic("host")
	}
}

func Server() courier.Transport { return server }

func Database() sqlx.DBExecutor { return database }

func Migrate() {
	err := migration.Migrate(database, nil)
	if err != nil {
		panic(err)
	}
}

func MigrateExprs() {
	err := migration.Migrate(database, os.Stdout)
	if err != nil {
		panic(err)
	}
}
