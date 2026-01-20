package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    mapCallbackFunction,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    mapbCallbackFunction,
		},
	}
}

func exitCallbackFunction(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCallbackFunction(*config) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Displays the names of 20 location areas\nmapb: Displays the previous 20 location areas")
	return nil
}

func mapCallbackFunction(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var locationAreas locationAreasResponse
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCallbackFunction(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.Previous

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var locationAreas locationAreasResponse
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	return nil
}