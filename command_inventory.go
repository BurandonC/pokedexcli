package main

import (
	"fmt"
)

// commandInventory displays the player's Pokeball inventory
func commandInventory(cfg *config, args ...string) error {
	if cfg.inventory == nil {
		return fmt.Errorf("inventory not initialized")
	}

	fmt.Println("Your Pokeball Inventory:")

	// Display only non-zero items
	balls := getAllPokeballs()
	hasItems := false

	// Define display order
	ballOrder := []string{"pokeball", "great", "ultra", "master", "net", "dusk"}

	for _, ballType := range ballOrder {
		count := cfg.inventory.Pokeballs[ballType]
		if count > 0 {
			ball := balls[ballType]
			fmt.Printf("  %s x%d\n", ball.DisplayName, count)
			hasItems = true
		}
	}

	if !hasItems {
		fmt.Println("  (empty)")
	}

	// Display total
	total := getTotalItems(cfg.inventory)
	fmt.Printf("\nTotal: %d Pokeballs\n", total)

	return nil
}
