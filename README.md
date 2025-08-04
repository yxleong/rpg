# Kage Missions

A fast-paced 2D RPG inspired by classic adventure games, blending smooth player movement, dynamic combat, and immersive world exploration.
Built with Go and Ebiten, this project demonstrates fundamental game development concepts such as animations, tilemaps, combat mechanics, and scene management.

## About the Game

Step into the shoes of a stealthy ninja on a quest through vibrant, pixel-art landscapes filled with challenges and secrets.
Master fluid controls, engage in strategic combat against enemies, and explore a richly designed world brought to life through animated sprites and detailed tilemaps.

## Progress & Current Status

- Implemented smooth player movement and collision detection
- Added enemy combat components with health and basic combat mechanics
- Developed player and enemy animations via spritesheets
- Created scene management system (Start, Pause, Game scenes with transitions)
- Integrated camera control and tilemap rendering from JSON data
- Refactored code for better structure and maintainability
- Core gameplay mechanics are functional
- Combat and health systems partially complete; attack logic in progress
- Web deployment planned but not yet implemented

## Upcoming Work

- Complete attack and health system
- Rebuild and enhance map and scene design
- Introduce player tasks and quests
- Finalize scene logic

---

## Running the Game

Make sure [Go is installed](https://golang.org/doc/install).

```bash
git clone https://github.com/yourusername/rpg-in-go.git
cd rpg-in-go
go run main.go
```

---

## Credits & Assets

- Player and environment assets from **Ninja Adventure - Asset Pack** (https://pixel-boy.itch.io/ninja-adventure-asset-pack)
- Map editing and tilemap creation using **Tiled** (https://www.mapeditor.org/)
- Tutorial series by [Coding with Sphere](https://www.youtube.com/playlist?list=PLvN4CrYN-8i7xnODFyCMty6ossz4eW0Cn) for the core game code
