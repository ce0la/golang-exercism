// Package weather shows what the current weather condition is in a given location.
package weather

// CurrentCondition stores the weather condition.
var CurrentCondition string
// CurrentLocation stores the location where the weather is being checked.
var CurrentLocation string

// Forecast returns what the weather condition is at a specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
