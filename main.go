package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		commands: commandRegistry(),
		mapOffset: -1,
	}
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
