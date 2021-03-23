package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Res struct {
	Global struct {
		NewConfirmed   int `json:"NewConfirmed" bson:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed" bson:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths" bson:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths" bson:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered" bson:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered" bson:"TotalRecovered"`
	} `json:"Global" bson:"Global"`
	Countries struct {
		Country        string `json:"Country" bson:"Country"`
		CountryCode    string `json:"CountryCode" bson:"CountryCode"`
		Date           string `json:"Date" bson:"Date"`
		NewConfirmed   int64  `json:"NewConfirmed" bson:"NewConfirmed"`
		NewDeaths      int64  `json:"NewDeaths" bson:"NewDeaths"`
		NewRecovered   int64  `json:"NewRecovered" bson:"NewRecovered"`
		Slug           string `json:"Slug" bson:"Slug"`
		TotalConfirmed int64  `json:"TotalConfirmed" bson:"TotalConfirmed"`
		TotalDeaths    int64  `json:"TotalDeaths" bson:"TotalDeaths"`
		TotalRecovered int64  `json:"TotalRecovered" bson:"TotalRecovered"`
	} `json:"Countries" bson:"Countries"`
	Date string `json:"Date" bson:"Date"`
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//fmt.Println("connected", err)
	collection := client.Database("firstdb").Collection("covid")
	fmt.Println("CLT", collection)
	response, err := http.Get("https://api.covid19api.com/summary")

	if err != nil {
		log.Printf("ERROR1 %s", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("DATA TYPE", reflect.TypeOf(data))
	//fmt.Println("RES11", string(data))

	var docs Res

	// Unmarshal the encoded JSON byte string into the slice
	//err1 := json.Unmarshal(data, &docs)
	err1 := bson.Unmarshal(data, &docs)
	if err1 != nil {
		fmt.Printf("ERROR2 %s", err1)
	}
	fmt.Println("\n 1DATA TYPE", reflect.TypeOf(err1))
	fmt.Println("Orig Data", err1)

	/*	for i := range data {

		// Put the document element in a new variable
		dat := data[i]

		fmt.Println("DAT", dat)
		// Call the InsertOne() method and pass the context and doc objects
		result, insertErr := collection.InsertOne(ctx, dat)
		fmt.Println("RESU", result)
		// Check for any insertion errors
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
		} else {
			fmt.Println("InsertOne() API result:", result)
		}
	}*/
}
