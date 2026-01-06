package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, caughtPokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", caughtPokemon.Pokemon.Name)
	}
	return nil
}
