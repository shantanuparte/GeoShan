package weathers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func GetWeatherInfo(place string) (*Location, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := buildUrl(place)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error in getting response: %v",err)
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

	return &data.Results[0], nil

}


func buildUrl(place string) string{
	encodePlace := url.QueryEscape(place)
	return fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", encodePlace)
}