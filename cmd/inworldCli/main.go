package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

func main() {
	fmt.Println("Please select table.")
	t := prompt.Input("> ", func(prompt.Document) []prompt.Suggest { return []prompt.Suggest{} })
	fmt.Println("You selected " + t)
}
