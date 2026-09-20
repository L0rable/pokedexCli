package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(strings.ToLower(scanner.Text()))
			if cmd, ok := getCommands()[input[0]]; ok {
				cmd.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
