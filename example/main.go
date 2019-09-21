package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/props/ini"
	_ "xffresk"
	"xffresk/infra"
)

func main() {
	conf := ini.NewIniFileCompositeConfigSource("example/config.ini")
	app := infra.New(conf)
	app.Start()

}
