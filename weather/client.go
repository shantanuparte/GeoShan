package weathers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

	if len(data.Results) == 0{
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
	if err != nil{
		return nil, fmt.Errorf("Error in respose: %v\n",err)
	}
	defer resp.Body.Close()

	
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Error in status: %v\n", resp.StatusCode)
	}
	
	
	//Unmarshalling
	var api ApiInfo
	err = json.NewDecoder(resp.Body).Decode(&api)
	if err != nil{
		return nil, fmt.Errorf("failed to decode: %v\n",err)
	}

	return &api, nil

}

func buildUrl(place string) string {
	encodePlace := url.QueryEscape(place)
	return fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", encodePlace)
}

func buildWeatherUrl(lat, lon float64) string {

	return fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current_weather=true", lat, lon)

}
