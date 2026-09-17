package main

import (
	"fmt"
	"os"

	"GeoShan/weather"

	"charm.land/lipgloss/v2"
)

func main() {
	fmt.Println("This is start of the project")
	place := os.Args[1]

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2)

	fmt.Println(style.Render("Hello, my first lipgloss project"))

	weathers.GetWeatherInfo(place)
}