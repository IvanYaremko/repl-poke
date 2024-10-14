package main

import (
	"fmt"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	locationId := args[0]
	response, err := cfg.pokeapiClient.LocationArea(locationId)
	if err != nil {
		return fmt.Errorf("command explore error: %w", err)
	}

	printEncounters(response.PokemoneEncounters)
	return nil
}

func printEncounters(encounters []pokeapi.Pokemon) {
	fmt.Println("Found pokemons:")
	for _, encounter := range encounters {
		fmt.Println(" 	- ", encounter.Pokemon.Name)
	}
}
