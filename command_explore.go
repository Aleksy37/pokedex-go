package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
	return errors.New("you must provide a location name")
	}

	locationName := args[0]
	encountersResp, err := cfg.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}
	
	fmt.Println("Found Pokemon:")
	for _, encounter := range encountersResp.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}
	return nil
}