package weathers


type Response struct{
	Results []Location `json:"results"`
}

// the weather info comes in 2 digit int format so adn i need to use a map to store and use it
type ApiInfo struct {
	Current struct {
		Temp          float32 `json:"temperature_2m"`
		WeatherInfo   uint8   `json:"weather_code"`
		Humidity      uint16  `json:"relative_humidity_2m"`
		Wind          float32 `json:"wind_speed_10m"`
		Precipitation float32 `json:"precipitation"`
	} `json:"current"`
}



type Location struct {
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
	Name    string  `json:"name"`
	Country string  `json:"country"`
}
