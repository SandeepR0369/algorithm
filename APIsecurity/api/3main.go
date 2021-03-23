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
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

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

func (docs Res) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.String, bsoncore.AppendString(nil, `fixed value`), nil
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	fmt.Println("connected", err)
	collection := client.Database("firstdb").Collection("covid")
	fmt.Println("CLT", collection)
	response, err := http.Get("https://api.covid19api.com/summary")

	if err != nil {
		log.Printf("ERROR1 %s", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("DATA TYPE", reflect.TypeOf(data))
	//fmt.Println("RES11", string(data))

	//var docs []interface{}

	// Unmarshal the encoded JSON byte string into the slice
	//err1 := json.Unmarshal(data, &docs)

	err1 := bson.MarshalBson(data, docs)
	if err1 != nil {
		fmt.Printf("ERROR2 %s", err1)
	}
	fmt.Println("\n 1DATA TYPE", reflect.TypeOf(err1))

	//var aplha []interface{}

	for i := range data {

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
	}
}

/**************2ND*****************/
/*
func main() {
	response, err := http.Get("https://api.covid19api.com/summary")

	if err != nil {
		log.Printf("ERROR1 %s", err)
	}

	//1/	data, _ := ioutil.ReadAll(response.Body)
	//1/	fmt.Println("RES11", string(data))
	var dat Res
	json.NewDecoder(response.Body).Decode(&Res)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//err = client.Ping(ctx, readpref.Primary())

	fmt.Println("connected", err)
	collection := client.Database("firstdb").Collection("covid")
	//fmt.Println("Collection type:", reflect.TypeOf(collection), "n")
	result, insertErr := collection.InsertOne(ctx, dat)
	json.NewEncoder(response)
	/**********************/

//docs := []byte(data)
/*
	// Unmarshal the encoded JSON byte string into the slice
	abc := json.Unmarshal(docs, &data)
	fmt.Println("RR1", abc)
	// Print MongoDB docs object type
	fmt.Println("nMongoFields Docs:", reflect.TypeOf(docs))

	// Iterate the slice of MongoDB struct docs
	for i := range docs {

		// Put the document element in a new variable
		doc := docs[i]

		fmt.Println("DOC", doc)
		// Call the InsertOne() method and pass the context and doc objects
		result, insertErr := collection.InsertOne(ctx, doc)
		fmt.Println("RESU", result)
		// Check for any insertion errors
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
		} else {
			fmt.Println("InsertOne() API result:", result)
		}
	}
	/**********************/
/*
}
*/
