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
	Name string `json:"name"`
	Url  string `json:"url"`
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
