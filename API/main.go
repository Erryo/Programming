package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	API_KEY  string = "b0fb6fd336130704cfbdb9e7360eca66"
	API_CALL string = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"
)

// var limiter *rate.Limiter = rate.NewLimiter(rate, 60)
var limiter *rate.Limiter = rate.NewLimiter(rate.Every(1*time.Minute), 30)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(city string) {
	url := fmt.Sprintf(API_CALL, city, API_KEY)

	if !limiter.Allow() {
		fmt.Println("Limit")
		return
	}
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("GET API_CALL failed:", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var cityData WeatherData
	err = json.Unmarshal(body, &cityData)
	if err != nil {
		fmt.Println("Error Unmarshaling body:", err)
		return
	}
	fmt.Printf("The current temperature in %s is %.2f°C\n", cityData.Name, cityData.Main.Temp)
}

func main() {
	getWeather("Ungheni")
}
