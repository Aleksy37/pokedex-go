package main


import (
	"time"
	"github.com/aleksy37/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.RespPokemonData{},
	}
	

	startRepl(cfg)
}

