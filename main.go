package main

import (
	weathers "GeoShan/weather"
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
)

func main() {
	place := os.Args[1]
	upper_place := strings.ToUpper(place)
	style := lipgloss.NewStyle().
		Bold(true).
		Width(60).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(2, 4).
		Align(lipgloss.Center)

	fmt.Println(style.Render(upper_place))
	
	var something *weathers.ApiInfo
	
	something, err := weathers.GetWeatherInfo(upper_place)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	
	fmt.Println(something)
	
}
