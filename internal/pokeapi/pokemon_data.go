package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	
)

func (c *Client) GetPokemonData(PokemonName string) (RespPokemonData, error) {
	url := baseURL + "/pokemon/" + PokemonName

	if entry, ok := c.cache.Get(url); ok {
		pokemonData := RespPokemonData{}
		err := json.Unmarshal(entry, &pokemonData)
		if err != nil {
			return RespPokemonData{}, err
		}
		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonData{}, err
	}

	
	pokemonData := RespPokemonData{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return RespPokemonData{}, err
	}
	
	c.cache.Add(url, dat)
	return pokemonData, nil
}
