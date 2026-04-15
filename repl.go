package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/aleksy37/pokedex-go/internal/pokeapi"
)
type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		
		cleanedInput := cleanInput(reader.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		
		commandName := cleanedInput[0]
		commandArgs := ""
		if len(cleanedInput) >= 2 {
			commandArgs = cleanedInput[1]
		}
		
		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(cfg, commandArgs)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
			} else {
				fmt.Printf("Unknown command: %s\n", commandName)
				continue
			}
			
		}
}
	
	
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of locations on the map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name: 		 "explore",
			description: "explore an area and return a list of encounters",
			callback: 	 commandExplore,
		},
	}
}