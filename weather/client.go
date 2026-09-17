package weathers

import (
	"fmt"
	"net/http"
	"time"
)

func GetWeatherInfo(place string) {

	fmt.Println("This funciton is running")
	client := &http.Client{
		Timeout: time.Second*5,
	}
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json",place)

	resp, err := client.Get(url)
	if err != nil{
		fmt.Println("Error in getting reponse")
		return
	}

	
}
