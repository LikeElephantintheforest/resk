package base

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {

	formatter := &prefixed.TextFormatter{}
	log.SetFormatter(formatter)

	//日志级别
	if os.Getenv("log.debug") == "true" {
		log.SetLevel(log.DebugLevel) //默认为info级别
	}

	formatter.ForceFormatting = true                //logrus-prefixed-formatter 支持
	formatter.SetColorScheme(&prefixed.ColorScheme{ //logrus-prefixed-formatter 支持
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		TimestampStyle:  "37", //日期打印的颜色
	})
	formatter.DisableColors = false //不禁用彩色日志
	formatter.ForceColors = true    //强行彩色化
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"

	//日志滚动策略
	//github.com/lestrrat/go-file-rotatelogs //todo 添加日志滚动策略
}
