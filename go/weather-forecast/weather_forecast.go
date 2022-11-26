// Package weather allows you to forecast the weather.
package weather

// CurrentCondition describes the current condition of the city.
var CurrentCondition string
// CurrentLocation describes the location of the city.
var CurrentLocation string

// Forecast the weather for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
