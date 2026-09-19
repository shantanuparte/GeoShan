package weathers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetLocatoinCoordinates(place string) (*Location, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := buildUrl(place)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error in getting response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	//Unmarshalling
	var data Response

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode")
	}

	if len(data.Results) == 0 {
		return nil, errors.New("location not found")
	}

	return &data.Results[0], nil

}

func GetWeatherInfo(info *Location) (*ApiInfo, error) {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := buildWeatherUrl(info.Lat, info.Lon)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error in respose: %w\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error in status: %v\n", resp.StatusCode)
	}

	//Unmarshalling
	var api ApiInfo
	err = json.NewDecoder(resp.Body).Decode(&api)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w\n", err)
	}

	return &api, nil

}

func GetLocatoinCoordinatesWithoutArguments() (*Location, error) {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := buildLocationUrl()
	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error: %w\n", err)

	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %v\n", resp.Status)
	}
	defer resp.Body.Close()

	//Unmarshalling
	var coordinates NoArg
	err = json.NewDecoder(resp.Body).Decode(&coordinates)

	if err != nil {
		return nil, fmt.Errorf("Error in unmarshalling: %w\n", err)
	}

	//Spliting through "," cause only one loc comes you forget it
	parts := strings.Split(coordinates.Loc, ",")

	if len(parts) != 2 {
		return nil, errors.New("Invalid location coordinates")
	}
	lat, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid longitude: %w", err)
	}

	lon, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid longitude: %w", err)
	}

	return &Location{
		Lat:     lat,
		Lon:     lon,
		Name:    coordinates.City,
		Country: coordinates.Con,
	}, nil

}

func buildUrl(place string) string {
	encodePlace := url.QueryEscape(place)
	return fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", encodePlace)
}

func buildWeatherUrl(lat, lon float64) string {

	return fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current=temperature_2m,weather_code,relative_humidity_2m,wind_speed_10m,precipitation", lat, lon)

}

func buildLocationUrl() string {
	return "https://ipinfo.io/json"
}

func WeahterCodeConvertion(code uint8) string {

	wmap := map[uint8]string{
		0:  "Clear sky",
		1:  "Mainly clear",
		2:  "Partly cloudy",
		3:  "Overcast",
		45: "Fog",
		48: "Depositing rime fog",
		51: "Light drizzle",
		53: "Moderate drizzle",
		55: "Dense intensity drizzle",
		56: "Light freezing drizzle",
		57: "Dense",
		61: "Slight rain",
		63: "Moderate rain",
		65: "Heavy rain",
		66: "Light freezing rain",
		67: "Heavy freezing rain",
		71: "Slight snowfall",
		73: "Moderate snowfall",
		75: "Heavy snowfall",
		77: "Snow grains",
		80: "Slight rain showers",
		81: "Moderate rain showers",
		82: "Violent rain showers",
		85: "Slight snow showers",
		86: "Heavy snow showers",
		95: "Thumderstorm",
		96: "Thunderstorm with slight hail",
		99: "Thunderstorm with heavy hail",
	}

	return wmap[code]
}
