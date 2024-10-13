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
	nextLocationUrl string
	prevLocationUrl string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if ok {
			err := command.callback(cfg)
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
	}
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}
