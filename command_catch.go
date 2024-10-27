package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if _, ok := cfg.caughtPokemon[name]; ok {
		return errors.New("pokemone is already caught")
	}

	pokemon, err := cfg.pokeapiClient.PokemonDetails(name)
	if err != nil {
		return fmt.Errorf("command catch response error: %w", err)
	}
	base := rand.IntN(pokemon.BaseExperience)
	target := float64(pokemon.BaseExperience) * 0.33
	fmt.Printf("throwing a ball at %s...\n", pokemon.Name)
	if base > int(target) {
		message := fmt.Sprintf("\U0001F386 \U0001F386 \U0001F386 caught: %s \U0001F386 \U0001F386 \U0001F386", pokemon.Name)
		fmt.Println(message)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		return nil
	} else {
		fmt.Println("failed to catch: ", pokemon.Name)
		return nil
	}
}
