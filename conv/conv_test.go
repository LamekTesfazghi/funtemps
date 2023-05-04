package conv

import (
	"reflect"
	"testing"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  string
	}

	tests := []test{
		{input: 134, want: "56.67"}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := FarhenheitToCelsius(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  string
	}

	tests := []test{
		{input: 10, want: "260.93"}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := FahrenheitToKelvin(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}


func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  string
	}

	tests := []test{
		{input: 20, want: "68.00"}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := CelsiusToFahrenheit(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  string
	}

	tests := []test{
		{input: 30, want: "303.15"}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := CelsiusToKelvin(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct { //lager en klasse
		input float64 //variabel for input med verdi	
		want  string
	}

	tests := []test{
		{input: 300, want: "26.85"}, //gir verdi vil ha verdi
	}

	for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
		got := KelvinToCelsius(tc.input) //got = input
		if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

	func TestKelvinToFahrenheit(t *testing.T) {
		type test struct { //lager en klasse
			input float64 //variabel for input med verdi	
			want  string
		}
	
		tests := []test{
			{input: 300, want: "80.33"}, //gir verdi vil ha verdi
		}
	
		for _, tc := range tests { //for loop som går gjennom alle verdiene til tests
			got := KelvinToFahrenheit(tc.input) //got = input
			if !reflect.DeepEqual(tc.want, got) { //sjekker om want er lik
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		}

}





