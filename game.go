package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"rpg-go/animations"
	"rpg-go/components"
	"rpg-go/constants"
	"rpg-go/entities"
	"rpg-go/spritesheet"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player            *entities.Player
	playerSpriteSheet *spritesheet.Spritesheet
	enemies           []*entities.Enemy
	potions           []*entities.Potion
	tilemapJSON       *TilemapJSON
	tilesets          []Tileset
	tilemapImage      *ebiten.Image
	cam               *Camera
	colliders         []image.Rectangle
}

func NewGame() *Game {
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImage, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	potionImage, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImage, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	tilesets, err := tilemapJSON.GenTilesets()
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSheet := spritesheet.NewSpriteSheet(4, 7, 16)

	return &Game{
		player: &entities.Player{
			Sprite: &entities.Sprite{
				Img: playerImage,
				X:   50.0,
				Y:   50.0,
			},
			Health: 3,
			Animations: map[entities.PlayerState]*animations.Animation{
				entities.Up:    animations.NewAnimation(5, 13, 4, 20.0),
				entities.Down:  animations.NewAnimation(4, 12, 4, 20),
				entities.Left:  animations.NewAnimation(6, 14, 4, 20.0),
				entities.Right: animations.NewAnimation(7, 15, 4, 20.0),
			},
			CombatComp: components.NewBasicCombat(3, 1),
		},
		playerSpriteSheet: playerSpriteSheet,
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Img: skeletonImage,
					X:   100.0,
					Y:   100.0,
				},
				FollowsPlayer: true,
				CombatComp:    components.NewEnemyCombat(3, 1, 30),
			},
			{
				Sprite: &entities.Sprite{
					Img: skeletonImage,
					X:   150.0,
					Y:   150.0,
				},
				FollowsPlayer: false,
				CombatComp:    components.NewEnemyCombat(3, 1, 30),
			},
		},
		potions: []*entities.Potion{
			{
				Sprite: &entities.Sprite{
					Img: potionImage,
					X:   210.0,
					Y:   100.0,
				},
				AmtHeal: 1.0,
			},
		},
		tilemapJSON:  tilemapJSON,
		tilemapImage: tilemapImage,
		tilesets:     tilesets,
		cam:          NewCamera(0.0, 0.0),
		colliders: []image.Rectangle{
			image.Rect(100, 100, 116, 116),
		},
	}
}

func (g *Game) Update() error {

	g.player.Dx, g.player.Dy = 0, 0

	// fixed speed
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.Dx = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.Dx = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Dy = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Dy = 2
	}

	g.player.X += g.player.Dx

	CheckCollisionHorizontal(g.player.Sprite, g.colliders)

	g.player.Y += g.player.Dy

	CheckCollisionVertical(g.player.Sprite, g.colliders)

	activeAnim := g.player.ActiveAnimation(int(g.player.Dx), int(g.player.Dy))
	if activeAnim != nil {
		activeAnim.Update()
	}

	for _, sprite := range g.enemies {

		sprite.Dx, sprite.Dy = 0.0, 0.0

		// accelerate: gradual speed
		if sprite.FollowsPlayer {
			if sprite.X < g.player.X {
				sprite.Dx += 1
			} else if sprite.X > g.player.X {
				sprite.Dx -= 1
			}
			if sprite.Y < g.player.Y {
				sprite.Dy += 1
			} else if sprite.Y > g.player.Y {
				sprite.Dy -= 1
			}
		}

		sprite.X += sprite.Dx
		CheckCollisionHorizontal(sprite.Sprite, g.colliders)

		sprite.Y += sprite.Dy
		CheckCollisionVertical(sprite.Sprite, g.colliders)
	}

	// for _, potion := range g.potions {
	// 	if g.player.X > potion.X {
	// 		g.player.Health += potion.AmtHeal
	// 		fmt.Printf("Picked up potion! Health: %d\n", g.player.Health)
	// 	}
	// }

	clicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	cX, cY := ebiten.CursorPosition()
	cX += int(g.cam.X)
	cY += int(g.cam.Y)
	g.player.CombatComp.Update()
	pRect := image.Rect(
		int(g.player.X),
		int(g.player.Y),
		int(g.player.X)+constants.Tilesize,
		int(g.player.Y)+constants.Tilesize,
	)

	deadEnemies := make(map[int]struct{})
	for index, enemy := range g.enemies {
		rect := image.Rect(
			int(enemy.X),
			int(enemy.Y),
			int(enemy.X)+constants.Tilesize,
			int(enemy.Y)+constants.Tilesize,
		)

		if rect.Overlaps(pRect) {
			if enemy.CombatComp.Attack() {
				g.player.CombatComp.Damage(enemy.CombatComp.AttackPower())
				if g.player.CombatComp.Health() <= 0 {
					fmt.Println("Player has been defeated")
					g.player.Health = 0
				} else {
					fmt.Printf("Player Health: %d\n", g.player.CombatComp.Health())
				}
			}
		}

		if cX > rect.Min.X && cX < rect.Max.X &&
			cY > rect.Min.Y && cY < rect.Max.Y {
			if clicked &&
				math.Sqrt(
					math.Pow(
						float64(cX)-g.player.X+(constants.Tilesize/2),
						2,
					)+math.Pow(
						float64(cY)-g.player.Y+(constants.Tilesize/2),
						2,
					),
				) < constants.Tilesize*5 {
				fmt.Println("damaging enemy")
				enemy.CombatComp.Damage(g.player.CombatComp.AttackPower())
			}
			if enemy.CombatComp.Health() <= 0 {
				deadEnemies[index] = struct{}{}
				fmt.Println("enemy has been eleminated")
			}
		}
	}

	if len(deadEnemies) > 0 {
		newEnemies := make([]*entities.Enemy, 0)
		for index, enemy := range g.enemies {
			if _, exists := deadEnemies[index]; !exists {
				newEnemies = append(newEnemies, enemy)
			}
		}
		g.enemies = newEnemies
	}

	g.cam.FollowsPlayer(g.player.X+8, g.player.Y+8, 320, 240)
	g.cam.Constrain(
		float64(g.tilemapJSON.Layers[0].Width)*16.0,
		float64(g.tilemapJSON.Layers[0].Height)*16.0,
		320,
		240,
	)

	return nil
}
