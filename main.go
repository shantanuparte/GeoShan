package main

import (
	weathers "GeoShan/weather"
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/common-nighthawk/go-figure"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Width(60).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(2, 4).
	Align(lipgloss.Center)

func main() {

	var location *weathers.Location
	var weather *weathers.ApiInfo
	var err error

	if len(os.Args) > 1 {
		location, err = weathers.GetLocatoinCoordinates(os.Args[1])
	} else {
		location, err = weathers.GetLocatoinCoordinatesWithoutArguments()
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	weather, err = weathers.GetWeatherInfo(location)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Render(location, weather)

}

func Render(location *weathers.Location, weather *weathers.ApiInfo) {

	fig := figure.NewFigure(location.Name, "doom", true)
	fig.Print()
	fmt.Println(location.Country)
	fmt.Printf("Temprature: %v\t Wind: %v\nHumidity: %v\tPrecipitation: %v\n", weather.Current.Temp, weather.Current.Wind, weather.Current.Humidity, weather.Current.Precipitation)
	fmt.Printf("   %v",weathers.WeahterCodeConvertion(weather.Current.WeatherInfo))
}
