package main

import (
	weathers "GeoShan/weather"
	"fmt"
	"os"
	"charm.land/lipgloss/v2"
)

func main() {

	if len(os.Args) > 1 {
		place := os.Args[1]

		var something *weathers.Location

		something, err := weathers.GetLocatoinCoordinates(place)
		if err != nil {

			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(something)

		var weather_info *weathers.ApiInfo

		weather_info, err = weathers.GetWeatherInfo(something)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("Country: %v\nCity: %v\nLatitude: %v\nLongitude: %v\n", something.Country, something.Name, something.Lat, something.Lon)

		fmt.Println(weather_info)
		fmt.Println(weathers.WeahterCodeConvertion(weather_info.Current.WeatherInfo))

	} else {
		var coordinates *weathers.Location

		coordinates, err := weathers.GetLocatoinCoordinatesWithoutArguments()
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		var weather_info_without_args *weathers.ApiInfo

		weather_info_without_args, err = weathers.GetWeatherInfo(coordinates)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		fmt.Printf("Country: %v\nCity: %v\nLatitude: %v\nLongitude: %v\n", coordinates.Country, coordinates.Name, coordinates.Lat, coordinates.Lon)
		fmt.Println(weather_info_without_args)

	}

	style := lipgloss.NewStyle().
		Bold(true).
		Width(60).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(2, 4).
		Align(lipgloss.Center)

	fmt.Println(style.Render("Currenty working"))
	fmt.Print("hii")

}
