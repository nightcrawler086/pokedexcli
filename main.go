package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Need to do an infinite for loop that prints a prompt
	// and accepts commands
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pokedex CLI")

	for {
		fmt.Print("pokedexcli> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
