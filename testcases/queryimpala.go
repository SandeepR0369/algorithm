package main

import (
       _ "github.com/alexbrainman/odbc"
        "github.com/jmoiron/sqlx"
        "fmt"
)

type Res struct {
	ID int `db:"id"`
        NAME string `db:"name"`
        DEPT string `db:"dept"`
}

func main() {
        conn, err := sqlx.Connect("odbc","DSN=IMPALA")
        if err != nil {
                fmt.Println("ERROR", err)
                panic("could not connect")
        }

        var res []Res
        query := "SELECT ID, NAME, DEPT FROM testdb.emp"
        if err := conn.Select(&res, query); err != nil {
                fmt.Println("ERROR", err)
                panic("error running query")
        }

        fmt.Println(res)
}
