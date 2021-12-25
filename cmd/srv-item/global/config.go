package global

import (
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

	App *appinfo.AppCtx
)

func init() {
	var config = &struct {
		Log         *conflogger.Log
		Server      *confhttp.Server
		DB          *confpostgres.Postgres
		ResPath     *string `env:""`
		MinioClient *confminioclient.MinioClient
	}{
		Server:      server,
		DB:          database,
		ResPath:     &ResPath,
		MinioClient: &MinioClient,
	}

	// pwd, _ := os.Getwd()

	App = appinfo.New(
		appinfo.WithName("srv-item"),
		// appinfo.WithMainRoot("."),
	)
	App.ConfP(config)

	confhttp.RegisterCheckerFromStruct(config)
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
