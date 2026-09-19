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

		fmt.Println(weather_info)

	}else {
		fmt.Println("upcoming feature ")

		
	}

		style := lipgloss.NewStyle().
		Bold(true).
		Width(60).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(2, 4).
		Align(lipgloss.Center)

		fmt.Println(style.Render("Kolhapur"))
		fmt.Print("hii")

}


