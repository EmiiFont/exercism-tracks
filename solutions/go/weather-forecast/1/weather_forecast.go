// Package weather contains tools to get the forecast for the city.
package weather

// CurrentCondition variable that holds the current forecast.
var CurrentCondition string

// CurrentLocation variable that holds the current location.
var CurrentLocation string

// Forecast calculates the weather condition of the city it takes city and contition parameters to get the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
