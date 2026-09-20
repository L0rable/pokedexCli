package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if (scanner.Scan()) {
			input := cleanInput(scanner.Text())
			if cmd, ok := getCommands()[input[0]]; ok {
				cmd.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
