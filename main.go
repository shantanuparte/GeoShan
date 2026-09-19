package main

import (
	weathers "GeoShan/weather"
	"fmt"
	"os"
	"github.com/common-nighthawk/go-figure"
	"charm.land/lipgloss/v2"
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

	
	Render(location,weather)
	
	
}

func Render(location *weathers.Location, weather *weathers.ApiInfo)(){
	
	fig := figure.NewFigure(location.Name, "doom",true)
	fig.Print()
	fmt.Println(style.Render(location.Name))
	fmt.Println(weather)
	fmt.Println(location)

}
