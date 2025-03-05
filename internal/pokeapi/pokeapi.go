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

func (c *Client) GetPokemonsByLocation(location string) ([]Pokemon, error) {
	var locationResp PokeApiLocationResponse

	url := c.baseUrl + "/location-area/" + location

	fmt.Println(url)

	if bodyBytes, ok := c.cache.Get(url); ok {
		// Already in cache
		err := json.Unmarshal(bodyBytes, &locationResp)
		if err != nil {
			return []Pokemon{}, nil
		}
	} else {

		// Create req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []Pokemon{}, nil
		}

		// Make request
		res, err := c.httpClient.Do(req)
		if err != nil {
			return []Pokemon{}, nil
		}
		defer res.Body.Close()

		// Check status
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
			return []Pokemon{}, errors.New("Response failed")
		}

		// Read bytes from body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return []Pokemon{}, nil
		}

		// As it was a miss, add new entry to cache
		c.cache.Add(url, bodyBytes)

		if err := json.Unmarshal(bodyBytes, &locationResp); err != nil {
			log.Fatal(err)
			return []Pokemon{}, nil
		}
	}

	return locationByNameToPokemonsList(locationResp), nil
}

func locationByNameToPokemonsList(locationResp PokeApiLocationResponse) []Pokemon {
	pokemons := []Pokemon{}
	for _, pokemonEncounter := range locationResp.PokemonEncounters {
		pokemons = append(pokemons, pokemonEncounter.Pokemon)
	}
	return pokemons
}
