package main

import (
	"github.com/alexbrainman/odbc"
)

func main() {
	//conn, _ := odbc.Connect("ODBC","DSN=impalacharter")
	conn, _ := sql.Open("ODBC","DSN=impalacharter")
	}
	fmt.Println("connected")
	conn.Close()
}
