package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	collection := cfg.caughtPokemon
	if len(collection) == 0 {
		return errors.New("no pokemon in collection")
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range collection {
		fmt.Println("	-", pokemon.Name)
	}
	return nil
}
