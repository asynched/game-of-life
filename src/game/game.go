package gol

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	GAME_WIDTH  = 640
	GAME_HEIGHT = 360
	GAME_TITLE  = "GAME OF LIFE"
	
	// Helper for setting the game scale
	GAME_SCALE  = 2

	// Helper for condicionally showing debug messages
	GAME_DEBUG  = true
)

type Game struct{
}

var grid[GAME_WIDTH][GAME_HEIGHT]bool

// Renders the grid to the screen
func (game *Game) drawGrid(screen *ebiten.Image) {
	for x := 0; x < GAME_WIDTH; x++ {
		for y := 0; y < GAME_HEIGHT; y++ {
			if grid[x][y] {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}
}

// Generates a new iteration of the game grid, evaluating the rules for the game of life.
func (game *Game) step(grid *[GAME_WIDTH][GAME_HEIGHT]bool) {
	var newGrid[GAME_WIDTH][GAME_HEIGHT]bool

	for x := 1; x < GAME_WIDTH - 1; x++ {
		for y := 1; y < GAME_HEIGHT - 1; y++ {
			cell := grid[x][y]

			top    := grid[x][y-1]
			right  := grid[x+1][y]
			bottom := grid[x][y+1]
			left   := grid[x-1][y]

			topRight    := grid[x+1][y-1]
			bottomRight := grid[x+1][y+1]
			bottomLeft  := grid[x-1][y+1]
			topLeft     := grid[x-1][y-1]

			cells := [8]bool{top, right, bottom, left, topRight, bottomRight, bottomLeft, topLeft}

			sum := 0

			for _, cell := range cells {
				if cell {
					sum++
				}
			}

			isNewCell := !cell && sum == 3
			doesCellSurvice := cell && (sum == 2 || sum == 3)

			if isNewCell || doesCellSurvice {
				newGrid[x][y] = true
			}
		}
	}

	*grid = newGrid
}

// Updates the screen on each render.
// Responsible for iterating for each step and drawing the game grid.
func (game *Game) Update(screen *ebiten.Image) error {
	game.step(&grid)
	game.drawGrid(screen)

	return nil
}

// Initializes the game layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return GAME_WIDTH, GAME_HEIGHT
}

// Initializes the game grid
func (game *Game) Setup() {
	seed := time.Now().UnixNano()

	if GAME_DEBUG {
		LogInfo("Generated random seed: ", seed)
	}

	rand.Seed(10)

	if GAME_DEBUG {
		LogInfo("Running game grid setup")
	}

	for x := 0; x < GAME_WIDTH; x++ {
		for y := 0; y < GAME_HEIGHT; y ++ {
			grid[x][y] = rand.Intn(4) == 0 
		}
	}
}

// Main execution function for the game of life
func Main() {
	game := Game{}
	game.Setup()

	if GAME_DEBUG {
		LogInfo("Setting window size and window title")
	}

	ebiten.SetWindowSize(GAME_WIDTH * GAME_SCALE, GAME_HEIGHT * GAME_SCALE)
	ebiten.SetWindowTitle(GAME_TITLE)
	
	if GAME_DEBUG {
		LogInfo("Starting game main loop")
	}

	err := ebiten.RunGame(&game)

	if err != nil {
		panic(err)
	}
}