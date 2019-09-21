package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	application := iris.Default()

	application.Get("/hello", func(context iris.Context) {
		context.WriteString("hello")
	})
	application.Get("/users/{id:uint64}", func(context iris.Context) {
		id := context.Params().GetUint64Default("id", 0)
		context.WriteString(strconv.Itoa(int(id)))
	})
	application.Get("/users/{action:path}", func(context iris.Context) {
		action := context.Params().Get("action")
		context.WriteString(action)
	})
	//添加路由
	v1Part := application.Party("/v1")
	//添加类似于Java中的Filter,middleware
	v1Part.Use(func(context iris.Context) {
		logrus.Info("自定义中间件/拦截器")
		context.Next() //like doChain
	})
	//通过前缀来区分
	v1Part.Get("/orders/{action:string prefix(a_)}", func(context context.Context) {
		paramTestPath := context.Params().Get("action")
		context.WriteString(paramTestPath)
	})
	//异常处理
	application.OnAnyErrorCode(func(context context.Context) {
		context.WriteString("系统繁忙")
	})
	//对应http状态码规定异常信息（会覆盖之前的异常处理）
	application.OnErrorCode(http.StatusNotFound, func(context context.Context) {
		context.WriteString("404 not found ，未找到相关资源")
	})

	err := application.Run(iris.Addr(":8082"))
	fmt.Println(err)

}
