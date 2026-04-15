package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationName string) error {
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