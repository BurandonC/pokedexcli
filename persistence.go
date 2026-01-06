package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/BurandonC/pokedexcli/internal/pokeapi"
)

// getPokedexDataPath returns the path to the saved Pokemon data file
// Creates the directory if it doesn't exist
func getPokedexDataPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pokedexDir := filepath.Join(homeDir, ".pokedexcli")

	// Create directory if it doesn't exist (0755 permissions)
	err = os.MkdirAll(pokedexDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(pokedexDir, "caught_pokemon.json"), nil
}

// savePokemon saves the caught Pokemon map to a JSON file
func savePokemon(caughtPokemon map[string]pokeapi.Pokemon) error {
	dataPath, err := getPokedexDataPath()
	if err != nil {
		return err
	}

	// Marshal to JSON with indentation for readability
	data, err := json.MarshalIndent(caughtPokemon, "", "  ")
	if err != nil {
		return err
	}

	// Write to file (0644 permissions)
	err = os.WriteFile(dataPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// loadPokemon loads caught Pokemon from the JSON file
// Returns empty map if file doesn't exist (not an error)
func loadPokemon() (map[string]pokeapi.Pokemon, error) {
	dataPath, err := getPokedexDataPath()
	if err != nil {
		return map[string]pokeapi.Pokemon{}, err
	}

	// Check if file exists
	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		// File doesn't exist - return empty map (not an error)
		return map[string]pokeapi.Pokemon{}, nil
	}

	// Read file
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return map[string]pokeapi.Pokemon{}, err
	}

	// Unmarshal JSON
	var caughtPokemon map[string]pokeapi.Pokemon
	err = json.Unmarshal(data, &caughtPokemon)
	if err != nil {
		return map[string]pokeapi.Pokemon{}, err
	}

	return caughtPokemon, nil
}
