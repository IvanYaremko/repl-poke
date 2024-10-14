package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("poke > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("invalid command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "list of available commands",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit application",
			callback:    calbackExit,
		},
		"mapf": {
			name:        "mapf",
			description: "go forward locations area of the world",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "go back locations area of the world",
			callback:    mapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "explore given area name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "attempt to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "inspects pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "see your pokemon collection",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}
