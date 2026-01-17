package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallbackFunction,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCallbackFunction,
		},
	}
}

func exitCallbackFunction() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCallbackFunction() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}