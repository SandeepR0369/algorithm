package main

import (
    "context"
    "flag"
    "sync"

    "github.com/moriarty-s3a/impalathing"
    log "github.com/sirupsen/logrus"
    "github.com/tsuna/gohbase"
    "github.com/tsuna/gohbase/hrpc"
)

var (
    hclient    gohbase.Client
    cluster    string
    impalaHost string
)

type Req struct {
    id     string
    values map[string]map[string][]byte
}

const (
    labZk  = "lab-cloudera-mstr-01.scope.charter.com,lab-cloudera-mstr-02.scope.charter.com,lab-cloudera-node-01.scope.charter.com"
    prodZk = "nce-cloudera-mstr-01.scope.charter.com,nce-cloudera-mstr-02.scope.charter.com,nce-cloudera-mstr-03.scope.charter.com,nce-cloudera-util-04.scope.charter.com,nce-cloudera-util-05.scope.charter.com"
)

func main() {
    chan1 := make(chan Req)
    var sendWG, receiveWG sync.WaitGroup
    var workerPool int

    var all, cmtsHub, nodeCmts, cmtsDivision bool
    flag.BoolVar(&all, "all", false, "Equivilient to setting all of the other flags to true")
    flag.BoolVar(&cmtsHub, "cmts-hub", false, "Attempt to populate the CMTS -> Hub relationship")
    flag.StringVar(&impalaHost, "impalaHost", "nce-lb-local.scope.charter.com", "The Impala hostname to use")
    flag.StringVar(&cluster, "cluster", "", "cluster you want data to push prod|lab")
    flag.IntVar(&workerPool, "workerPool", 10, "Customized workerpool size")
    flag.Parse()

    if cluster == "lab" {
        hclient = gohbase.NewClient(labZk)
        if hclient == nil {
            log.Fatal("Error connecting to lab cluster hbase")
        }
    } else if cluster == "prod" {
        hclient = gohbase.NewClient(prodZk)
        if hclient == nil {
            log.Fatal("Error connecting to prod cluster hbase")
        }
    }

    if all || cmtsHub {
        sendWG.Add(1)
        go getCmtsHubRelationship(&sendWG, chan1)
    }


    go func() {
        sendWG.Wait()
        close(chan1)
    }()
    counter := 0
    for i := 0; i < workerPool; i++ {
        receiveWG.Add(1)

        go func() {
            defer receiveWG.Done()
            for data := range chan1 {
                putReq, err := hrpc.NewPutStr(context.Background(), "entity-charter", data.id, data.values)

                if err != nil {
                    log.Errorf("Error creating put Request for row %+v \n", data.id)
                    continue
                }

                _, err = hclient.Put(putReq)

                if err != nil {
                    log.Errorf("Error putting to row %+v in entity-charter \n", putReq, err)
                }
                counter += 1
            }
        }()
    }
    receiveWG.Wait()
    log.Println("Total rows added:", counter)
}

func getCmtsHubRelationship(sendWG *sync.WaitGroup, chan1 chan Req) {
    defer sendWG.Done()
    conn, err := impalathing.Connect(impalaHost, "21000", impalathing.DefaultOptions)
    if err != nil {
        log.Fatalf("Error connecting to Impala %+v", err)
        return
    }
    query := "select cmts_id, hub_id from charterentity.cablemodem where cmts_id is not null and hub_id is not null group by cmts_id, hub_id order by cmts_id, count(*) desc "
    rows, err := conn.Query(query)
    if err != nil {
        log.Errorf("Error running Impala query %+v", err)
    }
    cmtsHub := make(map[string]string)
    var cmtsId, hubId, lastCmts string
    for rows.Next() {
        rows.Scan(&cmtsId, &hubId)
        log.Printf("Cmts: %s  Hub %s ", cmtsId, hubId)
        if cmtsId == lastCmts {
            continue
        }
        lastCmts = cmtsId
        cmtsHub[cmtsId] = hubId
    }
    log.Infof("Total cmtsHubCount found: %d \n", len(cmtsHub))
    for cmtsId, hubId := range cmtsHub {
        values := make(map[string]map[string][]byte)
        hubValues := make(map[string][]byte)
        hubValues["Derived_hub_hashKey"] = []byte(hubId)
        values["p"] = hubValues
        r := Req{id: cmtsId, values: values}
        chan1 <- r
    }
}