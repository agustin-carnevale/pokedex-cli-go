package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(url string) (PokeApiLocationsAreaResponse, error) {
	var locationsResp PokeApiLocationsAreaResponse

	if bodyBytes, ok := c.cache.Get(url); ok {
		// Already in cache
		err := json.Unmarshal(bodyBytes, &locationsResp)
		if err != nil {
			return PokeApiLocationsAreaResponse{}, err
		}
	} else {

		// Create req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokeApiLocationsAreaResponse{}, err
		}

		// Make request
		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokeApiLocationsAreaResponse{}, err
		}
		defer res.Body.Close()

		// Check status
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
			return PokeApiLocationsAreaResponse{}, errors.New("Response failed")
		}

		// Read bytes from body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return PokeApiLocationsAreaResponse{}, err
		}

		// As it was a miss, add new entry to cache
		c.cache.Add(url, bodyBytes)

		if err := json.Unmarshal(bodyBytes, &locationsResp); err != nil {
			log.Fatal(err)
			return PokeApiLocationsAreaResponse{}, err
		}

	}

	return locationsResp, nil
}

func (c *Client) GetPokemonsByLocation(location string) ([]PokemonItem, error) {
	var locationResp PokeApiLocationResponse

	url := c.baseUrl + "/location-area/" + location

	fmt.Println(url)

	if bodyBytes, ok := c.cache.Get(url); ok {
		// Already in cache
		err := json.Unmarshal(bodyBytes, &locationResp)
		if err != nil {
			return []PokemonItem{}, err
		}
	} else {

		// Create req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []PokemonItem{}, err
		}

		// Make request
		res, err := c.httpClient.Do(req)
		if err != nil {
			return []PokemonItem{}, err
		}
		defer res.Body.Close()

		// Check status
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
			return []PokemonItem{}, errors.New("Response failed")
		}

		// Read bytes from body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return []PokemonItem{}, err
		}

		// As it was a miss, add new entry to cache
		c.cache.Add(url, bodyBytes)

		if err := json.Unmarshal(bodyBytes, &locationResp); err != nil {
			log.Fatal(err)
			return []PokemonItem{}, err
		}
	}

	return locationByNameToPokemonsList(locationResp), nil
}

func locationByNameToPokemonsList(locationResp PokeApiLocationResponse) []PokemonItem {
	pokemons := []PokemonItem{}
	for _, pokemonEncounter := range locationResp.PokemonEncounters {
		pokemons = append(pokemons, pokemonEncounter.Pokemon)
	}
	return pokemons
}

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	var pokemonResp PokeApiPokemonResponse

	url := c.baseUrl + "/pokemon/" + pokemon

	fmt.Println(url)

	if bodyBytes, ok := c.cache.Get(url); ok {
		// Already in cache
		err := json.Unmarshal(bodyBytes, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
	} else {

		// Create req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		// Make request
		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		// Check status
		if res.StatusCode > 299 {
			// log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
			return Pokemon{}, errors.New("response failed")
		}

		// Read bytes from body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return Pokemon{}, err
		}

		// As it was a miss, add new entry to cache
		c.cache.Add(url, bodyBytes)

		if err := json.Unmarshal(bodyBytes, &pokemonResp); err != nil {
			log.Fatal(err)
			return Pokemon{}, err
		}
	}

	return parsePokemon(pokemonResp), nil
}

func parsePokemon(pokemonResp PokeApiPokemonResponse) Pokemon {
	return Pokemon{
		Name:           pokemonResp.Name,
		Height:         pokemonResp.Height,
		Weight:         pokemonResp.Weight,
		BaseExperience: pokemonResp.BaseExperience,
		Stats:          pokemonResp.Stats,
		Types:          pokemonResp.Types,
	}
}
