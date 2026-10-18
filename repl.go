package main

import (
	"errors"
	"fmt"
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
	commands  map[string]cliCommand
	client    internal.Client
	mapOffset int
	params    []string
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
		return errors.New("No location Id or Name was passed into explore")
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
