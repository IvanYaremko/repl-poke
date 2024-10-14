package main

import (
	"testing"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

func TestFormatMapLocation(t *testing.T) {
	testData := struct {
		input    pokeapi.RespLocationsAreaResults
		expected string
	}{
		input: pokeapi.RespLocationsAreaResults{
			Name: "location-1-name",
			URL:  "location-1-url",
		},

		expected: "Name: location-1-name",
	}

	actual := formatMapMessage(testData.input)
	if actual != testData.expected {
		t.Errorf("Actual not equal to expexted\nActual:	%s\nExpected: %s", actual, testData.expected)
	}
}
