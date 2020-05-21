package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userName := "root"
	password := "123456"
	dbName   := "gotest"
	dbHost   := "127.0.0.1:13306"
	db, _ := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			userName,
			password,
			dbHost,
			dbName))

	// insert
	ins, _ := db.Prepare("INSERT INTO students(name, age) values (?, ?)")
	res, _ := ins.Exec("qing", 20)
	id, _ := res.LastInsertId()
	fmt.Println(id)

	// update
	up,_ := db.Prepare("UPDATE students SET name=?, age=? WHERE id=?")
	res, _ = up.Exec("li", 30, id)

	// delete
	dl,_ := db.Prepare("DELETE FROM students WHERE id=?")
	res,_ = dl.Exec(id)

	// select
	sel, _  := db.Query("SELECT * FROM students ORDER BY id desc")
	//var id int
	for sel.Next()  {
		var name string
		var age int
		//name := ""
		//age := 0
		sel.Scan(&id, &name, &age)
		fmt.Println(id, name, age)
	}
}
