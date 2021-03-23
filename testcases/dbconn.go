package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() { 
	//ip := "10.10.7.114"
	ip := "24.24.32.101"
	
	db, err := sql.Open("mysql", fmt.Sprintf("oss:R3gioN0ss@tcp(%s:3306)/eng", ip))
	if err != nil {
 		fmt.Println(err) 
	} 	
	fmt.Println("connected")
	defer db.Close()
	err = db.Ping()
  	 if err != nil {
        panic(err.Error())
	}	 
}
