package main

import (
	"time"

	"github.com/BurandonC/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(15*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
