package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedexData:= cfg.pokedex
	if len(pokedexData) == 0 {
		return errors.New("you have not caught any pokemon, use catch <name> to catch some!")
	}
	fmt.Println("Your Pokedex:")
	for _, v := range pokedexData {
		fmt.Println("  -", v.Name)
	}
	return nil
}