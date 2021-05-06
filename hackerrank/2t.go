package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
 * Complete the 'getArticleTitles' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING author as parameter.
 *
 * URL for cut and paste:
 * https://jsonmock.hackerrank.com/api/articles?author=<authorName>&page=<num>
 *
 */

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
		StoryTitle  interface{} `json:"story_title"`
		StoryURL    interface{} `json:"story_url"`
		ParentID    interface{} `json:"parent_id"`
		CreatedAt   int         `json:"created_at"`
	} `json:"data"`
}

//Need to create new func which takes request and response writer

func getArticleTitles(author string) []string {

	query := "https://jsonmock.hackerrank.com/api/articles?"
	//fmt.Println(apiquery)

	resp, err := http.Get(query)
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp)
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
	//fmt.Println("AA", d.TotalPages)

	var i int
	for i = 1; i <= d.TotalPages; i++ {
		a := strconv.Itoa(i)
		apiquery := "https://jsonmock.hackerrank.com/api/articles?author=" + author + "&page=" + a

		resp1, err1 := http.Get(apiquery)
		if err1 != nil {
			panic(err1)
		}
		//fmt.Println(resp)
		body1, err1 := ioutil.ReadAll(resp1.Body)
		if err1 != nil {
			fmt.Println("cannot read body", err1)
		}
		//	fmt.Println(string(body))
		var dat BOOK

		err = json.Unmarshal(body1, &dat)
		if err1 != nil {
			fmt.Println("cannot unmarshal body", err)
		}
		for _, val := range dat.Data {
			fmt.Println("AS", val)
			title = append(title, val.Title)
		}

	}
	return title

}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	author := readLine(reader)

	result := getArticleTitles(author)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
