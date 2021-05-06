package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"strconv"
)

type BOOK struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Title       string      `json:"title"`
		URL         string      `json:"url"`
		Author      string      `json:"author"`
		NumComments int         `json:"num_comments"`
		StoryID     interface{} `json:"story_id"`
		StoryTitle  string      `json:"story_title"`
		StoryURL    interface{} `json:"story_url"`
		ParentID    interface{} `json:"parent_id"`
		CreatedAt   int         `json:"created_at"`
	} `json:"data"`
}

func main() {

	fmt.Println(jaffa("patricktomas"))

}

func jaffa(author string) []string {
	//  func jaffa(author string) []string {
	//apiquery := "https://jsonmock.hackerrank.com/api/articles?author=" + author

	query := "https://jsonmock.hackerrank.com/api/articles?"
	//fmt.Println(apiquery)

	resp, err := http.Get(query)
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read body", err)
	}
	//	fmt.Println(string(body))
	var d BOOK

	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("cannot unmarshal body", err)
	}

	//fmt.Println("D", d)

	var title []string
	//	fmt.Println("AA", d.TotalPages)

	//var i int
	for i := 1; i <= d.TotalPages; i++ {
		a := strconv.Itoa(i)
		apiquery := "https://jsonmock.hackerrank.com/api/articles?author=" + author + "&page=" + a
		fmt.Println(a, apiquery)
		resp1, err1 := http.Get(apiquery)
		if err1 != nil {
			panic(err1)
		}
		defer resp1.Body.Close()
		//fmt.Println(resp)
		body1, err1 := ioutil.ReadAll(resp1.Body)
		if err1 != nil {
			fmt.Println("cannot read body", err1)
		}
		//fmt.Println(string(body))
		var dat BOOK

		err = json.Unmarshal(body1, &dat)
		if err1 != nil {
			fmt.Println("cannot unmarshal body", err)
		}
		for _, val := range dat.Data {
			//fmt.Println(reflect.TypeOf(val.Title))
			title = append(title, val.Title)

			if val.Title == "" {
				title = append(title, val.StoryTitle)
			}
		}
	}
	return title
}

/*

"title": null,
"url": null,
"author": "patricktomas",
"num_comments": 376,
"story_id": null,
"story_title": "Steve Jobs has passed away.",
"story_url": "http://www.apple.com/stevejobs/",
"parent_id": null,
"created_at": 1317858143
}*/
