package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
	return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonData, err := cfg.pokeapiClient.GetPokemonData(pokemonName)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)

    catchRate := min(1.0, 100.0/float64(pokemonData.BaseExperience));

    if (rand.Float64() < float64(catchRate)) {
        fmt.Printf("%s was caught!\n", pokemonData.Name);
		cfg.pokedex[pokemonData.Name] = pokemonData
    } else {
        fmt.Printf("%s escaped!\n", pokemonData.Name);
    }
	return nil
}