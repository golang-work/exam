package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	userName := "root"
	password := "123456"
	dbName   := "gotest"
	dbHost   := "127.0.0.1:13306"
	// // 只有在实际操作数据库（查询、更新等）或调用 Ping 时才会真正连接
	db, _ := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			userName,
			password,
			dbHost,
			dbName))

	// db.SetMaxIdleConns(10) // 设置连接池中能够保持的最大空闲连接的数量。默认值是 2

	//fmt.Println("please exec show processlist")
	//time.Sleep(10 * time.Second)
	//fmt.Println("please exec show processlist again")
	//err := db.Ping()	// 进入 mysql 服务器通过“show processlist”命名可以看到 多了一个连接，Command 列是 Sleep
	//fmt.Println(err)
	//time.Sleep(10 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			db.Ping()
		}()
	}

	time.Sleep(10 * time.Second)
}
