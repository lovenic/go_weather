package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MainData struct {
	Temp      float32
	FeelsLike float32 `json:"feels_like"`
}
type WeatherData struct {
	Visibility int
	Main       MainData
}

func main() {
	fmt.Println("Hello weather")
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?units=metric&q=Minsk&appid=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Something went wrong with the weather request")
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	weatherData := WeatherData{}
	err = json.Unmarshal(data, &weatherData)

	fmt.Printf("Current temperature: %.1f | Feels like: %.1f", weatherData.Main.Temp, weatherData.Main.FeelsLike)
}
