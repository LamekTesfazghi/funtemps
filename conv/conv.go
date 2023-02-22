package conv

func FarhenheitToCelsius(value float64) float64 {
	return (value - 32) * (5.0 / 9)
}

func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}
func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}
func CelsiusToFarhenheit(value float64) float64 {
	return (value * (9.0 / 5)) + 32
}
func FarhenheitToKelvin(value float64) float64 {
	return (value-32)*(5/9) + 273.15
}
func KelvinToFarhenheit(value float64) float64 {
	return (value-273.15)*(9/5) + 32
}
