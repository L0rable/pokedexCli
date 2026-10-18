package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/L0rable/pokedexcli/internal"
	"github.com/L0rable/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		commands:      commandRegistry(),
		client:        internal.Client{},
		mapOffset:     -1,
		pokemonCaught: make(map[string]internal.PokemonApi),
	}
	// 5 minute interval
	config.client.Cache = *pokecache.NewCache(5 * time.Minute)

	for {
		if scanner.Err() != nil {
			fmt.Printf("scanner.Err(): %s\n", scanner.Err())
			return
		}

		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(strings.ToLower(scanner.Text()))
			if len(input) > 1 {
				config.params = input[1:]
			}

			if cmd, ok := commandRegistry()[input[0]]; ok {
				cmd.callback(&config)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
