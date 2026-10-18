package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/L0rable/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache pokecache.Cache
}

type LocationAreasApi struct {
	Result []LocationAreaApi `json:"results"`
}

type LocationAreaApi struct {
	Name              string `json:"name"`
	Url               string `json:"url"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreas(offSet int) (LocationAreasApi, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d", offSet*20)

	var data []byte
	if dataBytes, found := c.Cache.Get(url); found {
		data = dataBytes
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreasApi{}, err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasApi{}, err
		}

		// Add to cache
		c.Cache.Add(url, data)
	}

	locations := LocationAreasApi{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
		return LocationAreasApi{}, err
	}

	return locations, nil
}

func (c *Client) GetLocationArea(locationAreaIdOrName string) (LocationAreaApi, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationAreaIdOrName)

	var data []byte
	if dataBytes, found := c.Cache.Get(url); found {
		data = dataBytes
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaApi{}, err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaApi{}, err
		}

		// Add to cache
		c.Cache.Add(url, data)
	}

	location := LocationAreaApi{}
	err := json.Unmarshal(data, &location)
	if err != nil {
		fmt.Println(err)
		return LocationAreaApi{}, err
	}

	return location, nil
}

type PokemonApi struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetPokemon(pokemonIdOrName string) (PokemonApi, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonIdOrName)

	var data []byte
	if dataBytes, found := c.Cache.Get(url); found {
		data = dataBytes
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokemonApi{}, err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonApi{}, err
		}

		// Add to cache
		c.Cache.Add(url, data)
	}

	pokemon := PokemonApi{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println(err)
		return PokemonApi{}, err
	}

	return pokemon, nil
}
