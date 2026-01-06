package main

import (
	"fmt"
	"strings"

	"github.com/BurandonC/pokedexcli/internal/pokeapi"
)

// Pokeball represents a type of Pokeball with its properties
type Pokeball struct {
	Name           string
	DisplayName    string
	BaseMultiplier float64
	Rarity         string
}

// getAllPokeballs returns all available Pokeball types
func getAllPokeballs() map[string]Pokeball {
	return map[string]Pokeball{
		"pokeball": {
			Name:           "pokeball",
			DisplayName:    "Pokeball",
			BaseMultiplier: 1.0,
			Rarity:         "common",
		},
		"great": {
			Name:           "great",
			DisplayName:    "Great Ball",
			BaseMultiplier: 1.5,
			Rarity:         "uncommon",
		},
		"ultra": {
			Name:           "ultra",
			DisplayName:    "Ultra Ball",
			BaseMultiplier: 2.0,
			Rarity:         "rare",
		},
		"master": {
			Name:           "master",
			DisplayName:    "Master Ball",
			BaseMultiplier: 999.0, // Effectively infinite
			Rarity:         "ultra-rare",
		},
		"net": {
			Name:           "net",
			DisplayName:    "Net Ball",
			BaseMultiplier: 1.0, // Base, special calculated separately
			Rarity:         "uncommon",
		},
		"dusk": {
			Name:           "dusk",
			DisplayName:    "Dusk Ball",
			BaseMultiplier: 1.0, // Base, special calculated separately
			Rarity:         "uncommon",
		},
	}
}

// Ball name aliases
var ballAliases = map[string]string{
	"pokeball":   "pokeball",
	"poke":       "pokeball",
	"normal":     "pokeball",
	"greatball":  "great",
	"great":      "great",
	"gb":         "great",
	"ultraball":  "ultra",
	"ultra":      "ultra",
	"ub":         "ultra",
	"masterball": "master",
	"master":     "master",
	"mb":         "master",
	"netball":    "net",
	"net":        "net",
	"duskball":   "dusk",
	"dusk":       "dusk",
}

// getPokeball returns a Pokeball by name or alias
func getPokeball(name string) (*Pokeball, error) {
	normalizedName := strings.ToLower(name)

	// Check if it's an alias
	if canonicalName, exists := ballAliases[normalizedName]; exists {
		normalizedName = canonicalName
	}

	balls := getAllPokeballs()
	if ball, exists := balls[normalizedName]; exists {
		return &ball, nil
	}

	return nil, fmt.Errorf("unknown Pokeball type: %s", name)
}

// calculateCatchMultiplier calculates the final catch rate multiplier
func calculateCatchMultiplier(pokemon pokeapi.Pokemon, ballType string, cfg *config) float64 {
	ball, err := getPokeball(ballType)
	if err != nil {
		return 1.0
	}

	// Master Ball always catches
	if ball.Name == "master" {
		return 999.0
	}

	// Net Ball - 3x for Bug or Water types
	if ball.Name == "net" {
		if isPokemonType(pokemon, "bug") || isPokemonType(pokemon, "water") {
			return 3.0
		}
		return 1.0
	}

	// Dusk Ball - 3x for Ghost/Dark types or cave locations
	if ball.Name == "dusk" {
		isGhostOrDark := isPokemonType(pokemon, "ghost") || isPokemonType(pokemon, "dark")
		isCaveLocation := cfg.currentLocation != nil &&
			containsKeyword(*cfg.currentLocation, []string{"cave", "tunnel", "underground", "cavern"})

		if isGhostOrDark || isCaveLocation {
			return 3.0
		}
		return 1.0
	}

	// Standard balls use base multiplier
	return ball.BaseMultiplier
}

// isPokemonType checks if Pokemon has a specific type
func isPokemonType(pokemon pokeapi.Pokemon, typeToCheck string) bool {
	for _, t := range pokemon.Types {
		if strings.ToLower(t.Type.Name) == strings.ToLower(typeToCheck) {
			return true
		}
	}
	return false
}

// containsKeyword checks if a string contains any of the keywords
func containsKeyword(text string, keywords []string) bool {
	lowerText := strings.ToLower(text)
	for _, keyword := range keywords {
		if strings.Contains(lowerText, keyword) {
			return true
		}
	}
	return false
}
