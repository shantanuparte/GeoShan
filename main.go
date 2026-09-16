package main

import (
	"fmt"
	"charm.land/lipgloss/v2"
	weather/weathers
)

func main() {
	fmt.Println("This is start of the project")
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2)

	fmt.Println(style.Render("Hello, my first lipgloss project"))

	getWeatherInfo()

}
