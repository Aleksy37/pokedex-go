package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	
)

func (c *Client) ListPokemon(locationName string) (RespPokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName

	if entry, ok := c.cache.Get(url); ok {
		encountersResp := RespPokemonEncounters{}
		err := json.Unmarshal(entry, &encountersResp)
		if err != nil {
			return RespPokemonEncounters{}, err
		}
		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	
	encountersResp := RespPokemonEncounters{}
	err = json.Unmarshal(dat, &encountersResp)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	
	c.cache.Add(url, dat)
	return encountersResp, nil
}
