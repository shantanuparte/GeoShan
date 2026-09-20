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
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Location: %+v\n", location)

	weather, err = weathers.GetWeatherInfo(location)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	Render(location, weather)

}

func Render(location *weathers.Location, weather *weathers.ApiInfo) {

	fig := figure.NewFigure(location.Name, "doom", true)
	countryStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A7A7B3")).
		Italic(true).
		MarginTop(1).
		MarginBottom(1)

	label := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#D6D6E7"))

	value := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#D6D6E7"))

	temperature := label.Render("Temperature") +
		value.Render(fmt.Sprintf(": %v°C", weather.Current.Temp))

	wind := label.Render("Wind") +
		value.Render(fmt.Sprintf(": %v km/h", weather.Current.Wind))

	humidity := label.Render("Humidity") +
		value.Render(fmt.Sprintf(": %v%%", weather.Current.Humidity))

	precipitation := label.Render("Precipitation") +
		value.Render(fmt.Sprintf(": %v", weather.Current.Precipitation))

	condition := label.Render("Condition") +
		value.Render(fmt.Sprintf(": %v",
			weathers.WeahterCodeConvertion(weather.Current.WeatherInfo),
		))

	reanderBox := lipgloss.NewStyle().
		Width(50).
		Padding(1, 3).
		MarginTop(1).
		MarginBottom(1).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				temperature,
				wind,
				humidity,
				precipitation,
				condition,
			),
		)

	container := lipgloss.NewStyle().Padding(2, 4).Align(lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				fig.String(),
				countryStyle.Render(location.Country),
				reanderBox,
			),
		)
	fmt.Println(container)
}
