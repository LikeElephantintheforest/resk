package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {

	dsName := "root:123456@tcp(127.0.0.1:3306)/resk?charset=utf8&parseTime=true"
	db, e := sql.Open("mysql", dsName)

	if e != nil {
		fmt.Println(e)
	}

	//最大空闲连接数
	db.SetMaxIdleConns(2)
	//最大连接数
	db.SetMaxOpenConns(3)
	//最大存活时间
	db.SetConnMaxLifetime(7 * time.Hour)

	fmt.Print(db.Query("select now() "))

	defer db.Close()

}
