package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasApi struct {
	Result []LocationAreaApi `json:"results"`
}

type LocationAreaApi struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetLocationAreas(offSet int) (LocationAreasApi, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d", offSet*20)
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasApi{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasApi{}, err
	}

	locations := LocationAreasApi{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
		return LocationAreasApi{}, err
	}

	return locations, nil
}
