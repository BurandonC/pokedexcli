package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	// Validate arguments
	if len(args) < 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokemonName := args[0]

	// Determine ball type (default to pokeball)
	ballType := "pokeball"
	if len(args) > 1 {
		ballType = args[1]
	}

	// Validate ball type
	ball, err := getPokeball(ballType)
	if err != nil {
		return fmt.Errorf("unknown Pokeball type: %s", ballType)
	}

	// Check inventory
	if !hasItem(cfg.inventory, ball.Name, 1) {
		return fmt.Errorf("you don't have any %ss! Use 'bag' to check your inventory", ball.DisplayName)
	}

	// Check if already caught (no duplicates)
	if _, exists := cfg.caughtPokemon[pokemonName]; exists {
		return fmt.Errorf("you already have %s in your Pokedex!", pokemonName)
	}

	// Fetch Pokemon data from API
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Track attempts for this species
	if _, exists := cfg.catchAttempts[pokemon.Name]; !exists {
		cfg.catchAttempts[pokemon.Name] = 0
	}
	cfg.catchAttempts[pokemon.Name]++
	currentAttempts := cfg.catchAttempts[pokemon.Name]

	// Calculate catch rate
	multiplier := calculateCatchMultiplier(pokemon, ball.Name, cfg)

	// Master Ball always catches
	caught := false
	if ball.Name == "master" {
		caught = true
	} else {
		// Apply multiplier to reduce random range
		adjustedMax := int(float64(pokemon.BaseExperience) / multiplier)
		if adjustedMax < 1 {
			adjustedMax = 1
		}

		res := rand.Intn(adjustedMax)
		threshold := 40

		caught = res <= threshold
	}

	// Use the ball (consumed regardless of outcome)
	fmt.Printf("Throwing a %s at %s...\n", ball.DisplayName, pokemon.Name)
	err = useItem(cfg.inventory, ball.Name)
	if err != nil {
		return err
	}

	// Save inventory immediately
	if err := saveInventory(cfg.inventory); err != nil {
		fmt.Printf("Warning: Could not save inventory: %v\n", err)
	}

	// Handle outcome
	if !caught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	// Success!
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	// Create caught Pokemon with metadata
	caughtPokemon := newCaughtPokemon(pokemon, ball.Name, currentAttempts)
	cfg.caughtPokemon[pokemon.Name] = caughtPokemon

	// Save Pokemon data
	if err := savePokemon(cfg.caughtPokemon); err != nil {
		fmt.Printf("Warning: Could not save Pokemon: %v\n", err)
	} else {
		fmt.Println("Pokedex saved!")
	}

	return nil
}
