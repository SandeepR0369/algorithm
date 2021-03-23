package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v3"
)

func main() {
	var err error
	client, err := elastic.NewClient(elastic.SetURL("http://cloud.nyroc.rr.com:9200"))
	if err != nil {
		log.Fatalf("Error creating ES client %+v", err)
		return
	}

	//		searchresult, err := client.Search().Index("charter_entity").Query(elastic.NewTermQuery("_type", "CableModem")&&("enabled",1)).Size(1).Pretty(true).Do()

	Query_1 := elastic.NewTermQuery("_type", "CableModem")
	Query_2 := elastic.NewTermQuery("enabled", 1)

	newQuery := elastic.NewBoolQuery()
	newQuery = newQuery.Must(Query_1, Query_2)

	searchresult, err := client.Search().Index("charter_entity").Query(newQuery).Size(1).Pretty(true).Do()

	if searchresult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d hubs\n", searchresult.Hits.TotalHits)

		for _, hit := range searchresult.Hits.Hits {
			var f interface{}
			err := json.Unmarshal(*hit.Source, &f)
			if err != nil {
				fmt.Println("Cannot deserialize", err)
			}
			m := f.(map[string]interface{})
			for k, v := range m {
				switch vv := v.(type) {
				case string:
					fmt.Println(k, "is string", vv)
				case float64:
					fmt.Println(k, "is float64", vv)
				case []interface{}:
					fmt.Println(k, "is an array:")
					for i, u := range vv {
						fmt.Println(i, u)
					}
				default:
					fmt.Println(k, "is of a type I don't know how to handle")
				}
			}
		}

	} else {
		fmt.Println("No results found!!")
	}
}
