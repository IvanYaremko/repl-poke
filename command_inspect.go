package main

import (
	"fmt"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught %s", name)
	}

	printPokemonDetails(pokemon)
	return nil
}

func printPokemonDetails(pokemone pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemone.Name)
	fmt.Printf("Height: %d\n", pokemone.Height)
	fmt.Printf("Weight: %d\n", pokemone.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemone.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemone.Types {
		fmt.Println("	-", t.Type.Name)
	}
}
