package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"time"
	"xffresk/infra"
)

//dbx 数据库配置
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx infra.StarterContext) {
	//conf := ctx.Props()
	//数据库配置
	//err := kvs.Unmarshal(conf, &settings, "mysql")
	//if err != nil {
	//	panic(err)
	//}

	source := ctx.Props()

	settings := dbx.Settings{
		DriverName:      source.GetDefault("mysql.driveName", ""),
		User:            source.GetDefault("mysql.username", ""),
		Password:        source.GetDefault("mysql.password", ""),
		Host:            source.GetDefault("mysql.address", ""),
		MaxOpenConns:    5,
		MaxIdleConns:    2,
		ConnMaxLifetime: 7 * time.Hour,
		Options: map[string]string{
			"charset":   "utf8",
			"parseTime": "true",
		},
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName())
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}
