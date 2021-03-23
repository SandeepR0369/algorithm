package main

import (
        //"database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	//"flag"
)

type Res struct {
        Id string `db:"id"`
        Name    string `db:"name"`
	Host string `db:"host"`
}

func main() {
	//var drumHost string
	//flag.StringVar(&drumHost, "drum-host", "127.0.0.1", "Take a drum host to query")

        fmt.Println("started")
	drumDB, err := sqlx.Connect("mysql", "scope:sc0p3@tcp(150.181.64.40:3306)/config")
	//drumDB, err := sqlx.Connect("mysql", fmt.Sprintf("scope:sc0p3@tcp(%s:3306)/config", drumHost))
	fmt.Println(drumDB)
        fmt.Println("connected")
        if err != nil {
                fmt.Println("ERROR", err)
                panic("could not connect")
        }
	var res []Res
        query := "select id, name, host from config.object LIMIT 3"

        if err := drumDB.Select(&res, query); err != nil {
                fmt.Println("Failed to get configobject  %+v ", err)
                panic("error running query")
        }
        fmt.Println("connected configobject")
        fmt.Println(res)

}
