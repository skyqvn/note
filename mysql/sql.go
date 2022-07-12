package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Password  string
}

func main() {
	constr := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", constr)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := db.Exec("insert into users(name,password) value(?,?);", "sky", 2009917)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.LastInsertId())
	rows, err := db.Query("select id,name,password from users")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		user := new(User)
		if rows.Next() {
			err := rows.Scan(&user.Id, &user.Name, &user.Password)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			return
		}
		fmt.Printf("%d.\n%s\n%s\n",user.Id,user.Name,user.Password)
	}
}
