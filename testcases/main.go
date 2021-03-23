package main

import (
       _ "github.com/alexbrainman/odbc"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Res struct {
	Name string `db:"name"`
	Comment string `db:"comment"`
}

func main() {
        conn, err := sqlx.Connect("odbc","DSN=IMPALA")
	if err != nil {
		fmt.Println("ERROR", err)
		panic("could not connect")
	}

	var res []Res
	query := "SELECT name, comment FROM testdb.t1"
	if err := conn.Select(&res, query); err != nil {
		fmt.Println("ERROR", err)
		panic("error running query")
	}
	
	fmt.Println(res)
}	


