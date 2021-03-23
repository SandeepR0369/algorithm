package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type weatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {

	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/weather", getWeather)
	http.HandleFunc("/wind", getWind)

	http.ListenAndServe(":8081", nil)
}

func getHello(res http.ResponseWriter, t *http.Request) {
	res.Write([]byte("Hello World"))
}

func getWeather(res http.ResponseWriter, t *http.Request) {

	wd := getWeatherData()
	fmt.Println("!", wd)
	fmt.Println("Max", wd.Main.TempMax)
	fmt.Println("Min", wd.Main.TempMin)
}

func getWeatherData(res http.ResponseWriter, t *http.Request) weatherData {

	city := t.FormValue("q")
	fmt.Println(city)

	key := "fabc0f3ee2df47776dc03eed2998269f"

	query := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=" + key + "&units=Imperial"
	fmt.Println(query)

	resp, err := http.Get(query)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read body", err)
	}
	var wd weatherData

	err = json.Unmarshal(body, &wd)

	if err != nil {
		fmt.Printf("Issue withh the data %+v", err)
		return wd
	}

	return wd

}

func getWind(res http.ResponseWriter, t *http.Request) {

	wd := getWeatherData()
	fmt.Println("WindDeg", wd.Wind.Deg)
	fmt.Println("WindSpeed", wd.Wind.Speed)
}
