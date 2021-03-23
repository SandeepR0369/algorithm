package main

import (
    "log"
    "fmt"
    "time"
   // "github.com/koblas/impalathing"
     "github.com/moriarty-s3a/impalathing"
)

func main() {
    host := "nce-lb-local.scope.charter.com"
    port := 21000

    con, err := impalathing.Connect(host, port, impalathing.DefaultOptions)

    if err != nil {
        log.Fatal("Error connecting", err)
        return
    }

    query, err := con.Query("SELECT name, comment FROM testdb.t1")

    startTime := time.Now()
    total := 0
    for query.Next() {
        var (
            name     string
            comment      string
        //    yyyymm      int
        )

        query.Scan(&name, &acomment)
        total += 1

        fmt.Println(name, comment)
    }

    log.Printf("Fetch %d rows(s) in %.2fs", total, time.Duration(time.Since(startTime)).Seconds())
}
