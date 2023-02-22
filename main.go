package main

import (
	"flag"
	"fmt"

	conv "github.com/LamekTesfazghi/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cel float64
var kel float64
var out string
var funfacts string

func printTemperatures() {
	// ...
	fahrenheitFloat := 134.0
	celsius := conv.FarhenheitToCelsius(fahrenheitFloat)
	fmt.Printf("%2f °C\n", celsius)

	celsiusFloat := 56.67
	fahrenheit := conv.CelsiusToFarhenheit(celsiusFloat)
	fmt.Printf("%2f °F\n", fahrenheit)

	kelvinFloat := 26.85
	kelvin := conv.CelsiusToKelvin(kelvinFloat)
	fmt.Printf("%2f °K\n", kelvin)

	celsiusFloat = 300
	celsius = conv.KelvinToCelsius(celsiusFloat)
	fmt.Printf("%2f °C\n", celsius)

	kelvinFloat = 50
	kelvin = conv.FarhenheitToKelvin(kelvinFloat)
	fmt.Printf("%2f °K\n", kelvin)

	fahrenheitFloat = 283.15
	fahrenheit = conv.KelvinToFarhenheit(fahrenheitFloat)
	fmt.Printf("%2f °F\n", fahrenheit)
}

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()
	printTemperatures()

	if out == "C" && isFlagPassed("F") {
		fmt.Printf("Svaret blir %.2f", conv.FarhenheitToCelsius(fahr))
	} else if out == "F" && isFlagPassed("C") {
		fmt.Printf("Svaret blir %.2f", conv.CelsiusToFarhenheit(cel))
	} else if out == "K" && isFlagPassed("C") {
		fmt.Printf("Svaret blir %.2f", conv.CelsiusToKelvin(cel))
	} else if out == "C" && isFlagPassed("K") {
		fmt.Printf("Svaret blir %.2f", conv.KelvinToCelsius(kel))
	} else if out == "K" && isFlagPassed("F") {
		fmt.Printf("Svaret blir %.2f", conv.FarhenheitToKelvin(fahr))
	} else if out == "F" && isFlagPassed("K") {
		fmt.Printf("Svaret blir %.2f", conv.KelvinToFarhenheit(kel))
	} else {
		fmt.Println("Noe funket ikke")
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
