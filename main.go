package main

import (
	"fmt"
	"time"

	"github.com/BurandonC/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(15*time.Second, time.Minute*5)

	// Load saved Pokemon
	savedPokemon, err := loadPokemon()
	if err != nil {
		fmt.Printf("Warning: Could not load saved Pokemon: %v\n", err)
		savedPokemon = map[string]pokeapi.Pokemon{}
	}

	// Display welcome message
	if len(savedPokemon) > 0 {
		fmt.Printf("Welcome back! Loaded %d Pokemon from your Pokedex\n", len(savedPokemon))
	} else {
		fmt.Println("Welcome to Pokedex! Starting fresh")
	}

	cfg := &config{
		caughtPokemon: savedPokemon,
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
