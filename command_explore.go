package main

import (
	"errors"
	"fmt"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationId := args[0]
	response, err := cfg.pokeapiClient.LocationArea(locationId)
	if err != nil {
		return fmt.Errorf("command explore error: %w", err)
	}

	printEncounters(response.PokemoneEncounters)
	return nil
}

func printEncounters(encounters []pokeapi.Encounters) {
	fmt.Println("Found pokemons:")
	for _, encounter := range encounters {
		fmt.Println(" 	- ", encounter.Pokemon.Name)
	}
}
