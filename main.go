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
		commands:  commandRegistry(),
		client:    internal.Client{},
		mapOffset: -1,
	}
	// 5 minute interval
	config.client.Cache = *pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(strings.ToLower(scanner.Text()))
			if cmd, ok := commandRegistry()[input[0]]; ok {
				cmd.callback(&config)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
