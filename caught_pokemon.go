package main

import (
	"time"

	"github.com/BurandonC/pokedexcli/internal/pokeapi"
)

// CaughtPokemon represents a Pokemon that has been caught with metadata
type CaughtPokemon struct {
	Pokemon    pokeapi.Pokemon `json:"pokemon"`
	CaughtWith string          `json:"caught_with"`
	CaughtAt   time.Time       `json:"caught_at"`
	Attempts   int             `json:"attempts"`
}

// newCaughtPokemon creates a new CaughtPokemon instance
func newCaughtPokemon(pokemon pokeapi.Pokemon, ballType string, attempts int) CaughtPokemon {
	return CaughtPokemon{
		Pokemon:    pokemon,
		CaughtWith: ballType,
		CaughtAt:   time.Now(),
		Attempts:   attempts,
	}
}
