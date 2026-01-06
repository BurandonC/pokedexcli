package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	caughtPokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that Pokemon")
	}

	pokemon := caughtPokemon.Pokemon

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	// Display catch metadata
	fmt.Println()
	ball, _ := getPokeball(caughtPokemon.CaughtWith)
	fmt.Printf("Caught with: %s\n", ball.DisplayName)
	fmt.Printf("Caught on: %s\n", caughtPokemon.CaughtAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Attempts: %d\n", caughtPokemon.Attempts)

	return nil
}
