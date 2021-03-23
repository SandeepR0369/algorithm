package main

import (
       _ "github.com/alexbrainman/odbc"
	"database/sql"
	"fmt"
)

type Res struct {
	Jumbleseq string `db:"jumblesequence"`
	JumbleName string `db:"jumblename"`
}

func main() {
        conn, err := sqlx.Connect("odbc","DSN=impalacharter")
	if err != nil {
		panic("could not connect")
	}

	var res []Res
	query := "SELECT jumblesequence, jumblename FROM sandeep.jumblehash  limit 2"
	if err := conn.Select(&res, query); err != nil {
		panic("error running query")
	}
}	


