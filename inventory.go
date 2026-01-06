package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// Inventory represents the player's Pokeball inventory
type Inventory struct {
	Pokeballs map[string]int `json:"pokeballs"`
}

// getInventoryPath returns the path to the inventory file
func getInventoryPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pokedexDir := filepath.Join(homeDir, ".pokedexcli")
	err = os.MkdirAll(pokedexDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(pokedexDir, "inventory.json"), nil
}

// saveInventory saves inventory to JSON file
func saveInventory(inventory *Inventory) error {
	dataPath, err := getInventoryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataPath, data, 0644)
}

// loadInventory loads inventory from JSON file
func loadInventory() (*Inventory, error) {
	dataPath, err := getInventoryPath()
	if err != nil {
		return getDefaultInventory(), err
	}

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		return getDefaultInventory(), nil
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return getDefaultInventory(), err
	}

	var inventory Inventory
	err = json.Unmarshal(data, &inventory)
	if err != nil {
		return getDefaultInventory(), err
	}

	return &inventory, nil
}

// getDefaultInventory returns starting inventory
func getDefaultInventory() *Inventory {
	return &Inventory{
		Pokeballs: map[string]int{
			"pokeball": 10,
			"great":    3,
			"ultra":    0,
			"master":   0,
			"net":      0,
			"dusk":     0,
		},
	}
}

// hasItem checks if inventory has at least 'quantity' of an item
func hasItem(inventory *Inventory, ballType string, quantity int) bool {
	count, exists := inventory.Pokeballs[ballType]
	if !exists {
		return false
	}
	return count >= quantity
}

// useItem decrements an item from inventory
func useItem(inventory *Inventory, ballType string) error {
	if !hasItem(inventory, ballType, 1) {
		ball, _ := getPokeball(ballType)
		return fmt.Errorf("you don't have any %ss", ball.DisplayName)
	}

	inventory.Pokeballs[ballType]--
	return nil
}

// addItem adds items to inventory
func addItem(inventory *Inventory, ballType string, quantity int) {
	if _, exists := inventory.Pokeballs[ballType]; !exists {
		inventory.Pokeballs[ballType] = 0
	}
	inventory.Pokeballs[ballType] += quantity
}

// getTotalItems returns total count of all Pokeballs
func getTotalItems(inventory *Inventory) int {
	total := 0
	for _, count := range inventory.Pokeballs {
		total += count
	}
	return total
}

// generatePokeballRewards generates random Pokeball rewards for exploring
func generatePokeballRewards(locationName string) map[string]int {
	rewards := make(map[string]int)

	// 60% chance: 1-3 Pokeballs
	if rand.Intn(100) < 60 {
		rewards["pokeball"] = rand.Intn(3) + 1
	}

	// 25% chance: 1-2 Great Balls
	if rand.Intn(100) < 25 {
		rewards["great"] = rand.Intn(2) + 1
	}

	// 10% chance: 1 Ultra Ball
	if rand.Intn(100) < 10 {
		rewards["ultra"] = 1
	}

	// 3% base chance for specialty balls
	// Check for area bonuses first
	isCaveArea := containsKeyword(locationName, []string{"cave", "tunnel", "underground", "cavern"})
	isWaterArea := containsKeyword(locationName, []string{"water", "sea", "ocean", "lake", "river", "bay", "beach"})

	// Cave area: +20% chance for Dusk Ball
	if isCaveArea && rand.Intn(100) < 20 {
		rewards["dusk"] = 1
	}

	// Water area: +20% chance for Net Ball
	if isWaterArea && rand.Intn(100) < 20 {
		rewards["net"] = 1
	}

	// Base 3% for any specialty ball
	if rand.Intn(100) < 3 {
		specialBalls := []string{"net", "dusk"}
		chosen := specialBalls[rand.Intn(len(specialBalls))]

		// Don't double-reward if already got one from area bonus
		if _, exists := rewards[chosen]; !exists {
			rewards[chosen] = 1
		}
	}

	// 0.5% chance: Master Ball
	if rand.Intn(1000) < 5 {
		rewards["master"] = 1
	}

	return rewards
}
