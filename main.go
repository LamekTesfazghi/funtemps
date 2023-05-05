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

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene som brukeren skal bruke i kommandolinje
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit") //allokerer plass til fahr variabel med 64 bits
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "0.0", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse() //ser hva som er i kommandolinje og legger det inn i variabel 

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.
	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())
	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, kel, cel, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	//fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if isFlagPassed("F") && out == "C" {
		fmt.Println("F:", fahr, "til C:", conv.FarhenheitToCelsius(fahr)) // Kalle opp funksjonen FahrenheitToCelsius(fahr), som da skal returnere °C
	} else if isFlagPassed("F") && out == "K"{
		fmt.Println("F:", fahr, "til K:", conv.FahrenheitToKelvin(fahr))
	} else if isFlagPassed("C") && out == "F"{
		fmt.Println("C:", cel, "til F:", conv.CelsiusToFahrenheit(cel))
	} else if isFlagPassed("C") && out == "K"{
		fmt.Println("C:", cel, "til K:", conv.CelsiusToKelvin(cel))
	} else if isFlagPassed("K") && out == "C"{
		fmt.Println("K:", kel, "til C:", conv.KelvinToFahrenheit(kel))
	} else if isFlagPassed("K") && out == "F"{
		fmt.Println("K", kel, "til F:", conv.KelvinToFahrenheit(kel))
	} else {
		fmt.Println("Noe gikk galt")
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
