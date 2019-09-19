package main

import (
	"fmt"
	"github.com/tietang/props/ini"
)

func main() {
	conf := ini.NewIniFileConfigSource("example/config.ini")
	port := conf.GetIntDefault("app.server.port", 18000)
	fmt.Println(port)
}
