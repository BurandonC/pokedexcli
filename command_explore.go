package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	// Update current location for Dusk Ball bonus
	cfg.currentLocation = &location.Name

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	// Award random Pokeballs
	fmt.Println()
	rewards := generatePokeballRewards(location.Name)

	if len(rewards) > 0 {
		fmt.Println("Items found:")
		for ballType, quantity := range rewards {
			ball, _ := getPokeball(ballType)
			if quantity == 1 {
				fmt.Printf(" + %s\n", ball.DisplayName)
			} else {
				fmt.Printf(" + %s x%d\n", ball.DisplayName, quantity)
			}
			addItem(cfg.inventory, ballType, quantity)
		}

		// Save inventory
		if err := saveInventory(cfg.inventory); err != nil {
			fmt.Printf("Warning: Could not save inventory: %v\n", err)
		}
	} else {
		fmt.Println("No items found this time.")
	}

	return nil
}
