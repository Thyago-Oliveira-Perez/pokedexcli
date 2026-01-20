package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        	string
	description 	string
	callback     	func(*config) error
}

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		
		if len(words) == 0 {
			continue
		}
		
		commandName := words[0]
		if cmd, exists := commands[commandName]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", commandName)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}