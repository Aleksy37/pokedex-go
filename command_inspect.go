package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
	return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonData, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	
	fmt.Println("Name:", pokemonData.Name)
	fmt.Println("Height:", pokemonData.Height)
	fmt.Println("Weight:", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemonData.Stats {
		fmt.Printf("  -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemonData.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}

	return nil
}