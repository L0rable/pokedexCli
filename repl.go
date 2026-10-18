package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/L0rable/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands      map[string]cliCommand
	client        internal.Client
	mapOffset     int
	params        []string
	pokemonCaught map[string]internal.PokemonApi
}

func commandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of the 20 previous location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List all of the Pokemon located in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Chance to catch a specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display information about the caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for cmd := range c.commands {
		fmt.Printf("%s: %s\n", c.commands[cmd].name, c.commands[cmd].description)
	}
	return nil
}

func commandMap(c *config) error {
	c.mapOffset++

	locations, err := c.client.GetLocationAreas(c.mapOffset)
	if err != nil {
		return err
	}

	for _, l := range locations.Result {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapBack(c *config) error {
	if c.mapOffset > 0 {
		c.mapOffset--
	}

	locations, err := c.client.GetLocationAreas(c.mapOffset)
	if err != nil {
		return err
	}

	for _, l := range locations.Result {
		fmt.Println(l.Name)
	}
	return nil
}

func commandExplore(c *config) error {
	if len(c.params) < 1 {
		return errors.New("No Location Id or Name was passed into Explore command")
	}

	location, err := c.client.GetLocationArea(c.params[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring", location.Name, "...")
	fmt.Println("Found Pokemon:")
	for _, p := range location.PokemonEncounters {
		fmt.Println(" - ", p.Pokemon.Name)
	}

	return nil
}

const BaseExperience = 100

func commandCatch(c *config) error {
	if len(c.params) == 0 {
		return errors.New("No Pokemon Id or Name was passed into Catch command")
	}

	pokemon, err := c.client.GetPokemon(c.params[0])
	if err != nil {
		return err
	}

	// Check to see if Pokemon is captured
	roll := rand.Intn(BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if roll >= pokemon.BaseExperience {
		fmt.Println(pokemon.Name, "was caught!")
		c.pokemonCaught[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}

func commandInspect(c *config) error {
	if len(c.params) == 0 {
		return errors.New("No Pokemon Id or Name was passed into catch")
	}

	if pokemon, caught := c.pokemonCaught[c.params[0]]; caught {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

func commandPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokemonCaught {
		fmt.Println(" - ", pokemon.Name)
	}

	return nil
}
