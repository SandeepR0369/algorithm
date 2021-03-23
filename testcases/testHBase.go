package main

import (
        "flag"
        "log"

        "go.scope.charter.com/hbase"
        "go.scope.charter.com/hbase/pb"
)

func main() {
        flag.Parse()
        hclient := hbase.NewClient("entity-dev")
        rows := make([]*pb.CellSet_Row, 0)
        key := []byte("402a2dafea06f85953ffb89838fd75dc6b8335ea")
        values := make([]*pb.Cell, 0)
        values = append(values, &pb.Cell{
                Row:    key,
                Column: []byte("p:Eric_test"),
                Data:   []byte("asdf"),
        })
        row := &pb.CellSet_Row{
                Key:    key,
                Values: values,
        }
        rows = append(rows, row)
        if err := hclient.Put(&pb.CellSet{Rows: rows}); err != nil {
                log.Printf("Failed to put: %+v", err)
        }
}

