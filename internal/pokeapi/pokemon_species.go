package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetSpecies(pokemonName string) (PokemonSpecies, error) {
	url := baseURL + "/pokemon-species/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		speciesResp := PokemonSpecies{}
		err := json.Unmarshal(val, &speciesResp)
		if err != nil {
			return PokemonSpecies{}, err
		}

		return speciesResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonSpecies{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonSpecies{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonSpecies{}, err
	}

	speciesResp := PokemonSpecies{}
	err = json.Unmarshal(dat, &speciesResp)
	if err != nil {
		return PokemonSpecies{}, err
	}

	c.cache.Add(url, dat)
	return speciesResp, nil
}
