package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		fmt.Println("echoing: ", text)

		commands := getCommands()

		switch text {
		case "help":
			fmt.Println(commands[text].name)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "help commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit application",
			callback: func() error {
				os.Exit(0)
				return nil
			},
		},
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are your available commands:")
	fmt.Println("- help")
	fmt.Println("- exit")
	fmt.Println("")
	return nil
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}
