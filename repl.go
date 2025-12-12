package main

import "strings"

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}
	return words
}