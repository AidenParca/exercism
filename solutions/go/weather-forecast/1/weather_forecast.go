//Package weather will forcast the weather based on your city and conditions.
package weather
//CurrentCondition will be related to your Current Condition.
var CurrentCondition string
//CurrentLocation will be related to your Current Location.
var CurrentLocation string
//Forecast function will return current weather condition based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
