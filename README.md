# **Pokedex CLI**  
*A Go-based Pokédex with RPG elements—catch, battle, and level up Pokémon in your terminal!*  

![Go](https://img.shields.io/badge/Go-1.20%2B-blue)  
![License](https://img.shields.io/badge/License-MIT-green)  
![Status](https://img.shields.io/badge/Status-Alpha-yellow)  

Fetch Pokémon data, simulate battles, and build your party—all from the command line.  

---

## **✨ Features**  
✅ **Core**  
- Fetch Pokémon stats, types, and abilities via <a href="https://pokeapi.co" target="_blank" rel="noopener noreferrer">PokeAPI</a>.  
- Lightweight and fast (no dependencies beyond Go).  

🔜 **Future** *(see [Roadmap](#-roadmap))*  
- Battle simulations, random encounters, and evolution.  
- Save progress between sessions.  

---

## **📦 Installation**  
```bash
git clone https://github.com/BurandonC/pokedexcli.git
cd pokedexcli
go build -o pokedex cmd/main.go
./pokedex --name pikachu

# Fetch Pokémon
./pokedex 
Pokedex > <command>

# Battle (coming soon!)
./pokedex battle <Pokemon 1> <Pokemon 2>
```

## 🚧 Roadmap

### Quality of Life
- (🔴 = not started, 🟡 = in progress, 🟢 = completed)
- 🔴 ⌨️ **Command History**: Cycle through past commands with arrow keys.
- 🔴 🧪 **Testing**: Expand unit tests and refactor for testability.

### Gameplay
- 🔴 ⚔️ **Battles**: Simulate turn-based fights between Pokémon.
- 🔴 🌿 **Random Encounters**: Wild Pokémon appear as you explore.
- 🔴 🎮 **Exploration**: Interactive area navigation (e.g., "go left").
- 🔴 🏆 **Party System**: Catch Pokémon, level them up, and form a team.
- 🔴 ⏳ **Evolution**: Pokémon evolve after meeting conditions.
- 🔴 🔴 **Pokéballs**: Different catch rates (Pokéball vs. Ultra Ball).

### Persistence
- 🟢 💾 **Save Progress**: Store your Pokédex and party on disk.

---
## **🤝 Contributing**

PRs welcome! Discuss ideas in <a href="https://github.com/BurandonC/pokedexcli/issues" target="_blank" rel="noopener noreferrer">Issues</a>.
