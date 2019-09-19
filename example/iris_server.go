package main

import (
	"fmt"
	"github.com/kataras/iris"
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

	err := application.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
