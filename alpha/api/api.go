package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//	"github.com/olivere/elastic"
	elastic "gopkg.in/olivere/elastic.v7"
)

type Data struct {
	MetaData struct {
		OneInformation     string `json:"1. Information"`
		TwoSymbol          string `json:"2. Symbol"`
		ThreeLastRefreshed string `json:"3. Last Refreshed"`
		FourInterval       string `json:"4. Interval"`
		FiveOutputSize     string `json:"5. Output Size"`
		SixTimeZone        string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries5Min map[string]Attributes `json:"Time Series (5min)"`
}
type Attributes struct {
	OneOpen    string `json:"1. open"`
	TwoHigh    string `json:"2. high"`
	ThreeLow   string `json:"3. low"`
	FourClose  string `json:"4. close"`
	FiveVolume string `json:"5. volume"`
}

func main() {
	response, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=BMJ0NM2DZIU3GYJG")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(string(responseData))
	var data Data
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		log.Println("cannot unmarshal data", err)
	}
	///fmt.Println(data)
	/***********************************************************/

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	dataJSON, err := json.Marshal(data)
	//dataJSON, err := json.Marshal(newStudent)

	js := string(dataJSON)
	_, err = esclient.Index().
		Index("trading").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		//	panic(err)
		fmt.Println("!@", err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

	/*&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&*/

	/*	searchSource := elastic.NewSearchSource()
		searchSource.Query(elastic.NewMatchQuery("name", "Doe"))

		/* this block will basically print out the es query */
	//	queryStr, err1 := searchSource.Source()
	//	queryJs, err2 := json.Marshal(queryStr)

	/*	if err1 != nil || err2 != nil {
			fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
		}
		fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
		/* until this block */

	/*	searchService := esclient.Search().Index("trading").SearchSource(searchSource)

		searchResult, err := searchService.Do(ctx)
		if err != nil {
			fmt.Println("[ProductsES][GetPIds]Error=", err)
			return
		}

		for _, hit := range searchResult.Hits.Hits {
			var student Data
			err := json.Unmarshal(hit.Source, &student)
			if err != nil {
				fmt.Println("[Getting Students][Unmarshal] Err=", err)
			}

			students = append(students, student)
		}

		if err != nil {
			fmt.Println("Fetching student fail: ", err)
		} else {
			for _, s := range students {
				fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
			}
		}
	*/
}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

/*
func Jaffa() {

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	dataJSON, err := json.Marshal(data)
	//dataJSON, err := json.Marshal(newStudent)

	js := string(dataJSON)
	_, err = esclient.Index().
		Index("trading").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

	var students []Data
	//{***(0)********}
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "Doe"))

	/* this block will basically print out the es query */
/*	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

/*	searchService := esclient.Search().Index("trading").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student Data
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}

}*/
