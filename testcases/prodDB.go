package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//type Res struct {
//	FkSystem string `db:"fkSystem"`
//	HubId    string `db:"hubID"`
//}

type Res struct {
        PkSystem string `db:"pkSystem"`
        ShortName    string `db:"shortName"`
}

func main() {

	//ip := "10.10.7.114"
	ip := "24.24.32.116"	
	//ip := "172.21.26.192"


	repDB, err := sqlx.Connect("mysql", fmt.Sprintf("oss:R3gioN0ss@tcp(%s:3306)/eng", ip))
	fmt.Println("connected")
	fmt.Println(repDB)
	if err != nil {
		fmt.Println("ERROR", err)
		panic("could not connect")
	}

	var res []Res
	//query := "SELECT fkSystem, hubID FROM eng.Hub LIMIT 3"
	query := "SELECT pkSystem, shortName FROM eng.System LIMIT 3"


	if err := repDB.Select(&res, query); err != nil {
		fmt.Println("Failed to get altHubID %+v ", err)
		panic("error running query")
	}
	fmt.Println("connected rhubs")
	fmt.Println(res)
}
