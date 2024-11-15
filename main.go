package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

type GeoLocationResponse struct {
	Loc string `json:"loc"`
}

func getWeatherData(latitude, longitude string) (WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", latitude, longitude)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil
}

func getUserLocation() (string, error) {
	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var geoData GeoLocationResponse
	err = json.Unmarshal(body, &geoData)
	if err != nil {
		return "", err
	}

	return geoData.Loc, nil
}

func determineGrassStatus(temperature, windSpeed float64) string {
	if temperature < 0 { // 0 = 32 degrees (idk why i didnt just use f for this lmao)
		return "The grass is currently frozen, so you are safe to go anywhere outside."
	} else if windSpeed > 50 {
		return "The wind speeds are high, so you should be pushed away from grass. It is safe to go outside."
	}
	return "It's a nice day, so do not go outside, no matter the reason."
}

func main() {
	a := app.New()
	w := a.NewWindow("Anti Grass Weather App")
	w.Resize(fyne.NewSize(400, 300))

	resultLabel := widget.NewLabel("Loading weather...")

	location, err := getUserLocation()
	if err != nil {
		resultLabel.SetText("Error fetching location")
	} else {
		parts := strings.Split(location, ",")
		if len(parts) != 2 {
			resultLabel.SetText("Invalid location format")
		} else {
			latitude := strings.TrimSpace(parts[0])
			longitude := strings.TrimSpace(parts[1])

			weather, err := getWeatherData(latitude, longitude)
			if err != nil {
				resultLabel.SetText("Error fetching weather data")
			} else {
				temperatureF := (weather.CurrentWeather.Temperature * 9 / 5) + 32

				resultText := fmt.Sprintf("Temperature: %.2f°F\nWind Speed: %.2f km/h\n", temperatureF, weather.CurrentWeather.WindSpeed)
				grassStatus := determineGrassStatus(weather.CurrentWeather.Temperature, weather.CurrentWeather.WindSpeed)
				resultLabel.SetText(resultText + grassStatus)
			}
		}
	}

	w.SetContent(container.NewVBox(
		widget.NewLabel("Anti Grass Weather App"),
		resultLabel,
	))

	w.ShowAndRun()
}
