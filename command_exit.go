package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	// Save Pokemon before exiting
	if err := savePokemon(cfg.caughtPokemon); err != nil {
		fmt.Printf("Warning: Could not save Pokemon: %v\n", err)
	} else {
		fmt.Println("Pokedex saved!")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
