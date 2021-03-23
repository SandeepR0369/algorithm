package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type auto struct {
	Predictions []struct {
		Description       string `json:"description"`
		MatchedSubstrings []struct {
			Length int `json:"length"`
			Offset int `json:"offset"`
		} `json:"matched_substrings"`
		PlaceID              string `json:"place_id"`
		Reference            string `json:"reference"`
		StructuredFormatting struct {
			MainText                  string `json:"main_text"`
			MainTextMatchedSubstrings []struct {
				Length int `json:"length"`
				Offset int `json:"offset"`
			} `json:"main_text_matched_substrings"`
			SecondaryText string `json:"secondary_text"`
		} `json:"structured_formatting"`
		Terms []struct {
			Offset int    `json:"offset"`
			Value  string `json:"value"`
		} `json:"terms"`
		Types []string `json:"types"`
	} `json:"predictions"`
	Status string `json:"status"`
}

func main() {

	http.HandleFunc("/output", autocomplete)
	http.ListenAndServe(":8082", nil)
}

func autocomplete(w http.ResponseWriter, t *http.Request) {

	input := t.FormValue("input")
	fmt.Println(input)

	key := "AIzaSyDPDpxSyGaRaUZQcMsslrS1agul3cIVh1I"

	query := "https://maps.googleapis.com/maps/api/place/autocomplete/json?input=" + input + "&key=" + key

	fmt.Println(query)

	resp, err := http.Get(query)

	if err != nil {
		fmt.Println("Issue with api url", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read body", err)
	}

	//	fmt.Println("RESP", string(body))
	var ac auto

	err = json.Unmarshal(body, &ac)

	if err != nil {
		fmt.Printf("Issue withh the data %+v", err)
		return
	}

	fmt.Println("A", ac)

}
