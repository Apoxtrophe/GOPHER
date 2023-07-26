// Non-Reusable code to be placed in main.go

package main

import (
	//"log"
	"math"

	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewGame() *Game {
	g := &Game{}
	g.arr1 = make([][]int, screenHeight/pixelSize)
	g.arr2 = make([][]int, screenHeight/pixelSize)
	g.pixels = make([]byte, screenWidth*screenHeight*4)
	for i := range g.arr1 {
		g.arr1[i] = make([]int, screenWidth/pixelSize)
		g.arr2[i] = make([]int, screenWidth/pixelSize)
	}
	return g
}

func (g *Game) DrawPixels(screen *ebiten.Image) {
	for row := 0; row < len(g.arr1); row++ {
		for col := 0; col < len(g.arr1[row]); col++ {
			clr := ElementMap[g.arr1[row][col]].Color
			for y := 0; y < pixelSize; y++ {
				for x := 0; x < pixelSize; x++ {
					i := ((row*pixelSize+y)*screenWidth + (col*pixelSize + x)) * 4
					g.pixels[i+0] = clr.R
					g.pixels[i+1] = clr.G
					g.pixels[i+2] = clr.B
					g.pixels[i+3] = clr.A
				}
			}
		}
	}
	screen.WritePixels(g.pixels)
}

func (g *Game) UpdateCells() {
	alive := g.AliveCells_Shuffled()
	g.reset_arr2()              //Clears arr2
	for _, pos := range alive { //Delete errant cells from the borders of the screen
		row, col := pos[0], pos[1]
		if row <= 1 || col <= 1 || row >= len(g.arr1)-2 || col >= len(g.arr1[0])-2 {
			g.arr1[row][col] = 0
			continue
		}
		switch g.arr1[row][col] {

		default:
			g.arr2[row][col] = g.arr1[row][col]
		}
	}
	g.arr1, g.arr2 = g.arr2, g.arr1 //Swap arrays
}

var prevMouseX, prevMouseY int

func (g *Game) Input() {
	x, y := ebiten.CursorPosition()

	//Clamp to world bounds
	world_x := clamp(x/pixelSize, 0, len(g.arr1[0])-1)
	world_y := clamp(y/pixelSize, 0, len(g.arr1)-1)
	mouse_one := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	mouse_two := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

	// Brush Sizing
	_, wheelY := ebiten.Wheel()
	if wheelY > 0 {
		g.brushSize++
	} else if wheelY < 0 {
		g.brushSize--
	}
	if g.brushSize < 1 {
		g.brushSize = 1
	}
	if g.brushSize > 100 {
		g.brushSize = 100
	}
	radius := float64(g.brushSize) / 2.0
	// Clicking detection
	if mouse_one || mouse_two {
		dx := float64(world_x - prevMouseX)
		dy := float64(world_y - prevMouseY)
		length := math.Sqrt(float64(dx*dx + dy*dy))
		if length > 0 {
			dx /= (length)
			dy /= (length)
		}
		for i := 0; i <= int(length); i++ {
			x := prevMouseX + int(float64(i)*dx)
			y := prevMouseY + int(float64(i)*dy)
			for row := -radius; row <= radius; row++ {
				for col := -radius; col <= radius; col++ {
					dist := math.Hypot(float64(row), float64(col))
					if dist <= radius {
						ix := clamp(x+int(col), 0, len(g.arr1[0])-1)
						iy := clamp(y+int(row), 0, len(g.arr1)-1)
						if mouse_one {
							g.arr1[iy][ix] = g.index
						} else {
							g.arr1[iy][ix] = 0
						}
					}
				}
			}
		}
	}
	prevMouseX = world_x
	prevMouseY = world_y
}

func (g *Game) DEBUG_INFO(screen ebiten.Image, debugMode bool) {
	if debugMode {
		ebitenutil.DebugPrint(&screen, fmt.Sprintf("TPS: %.2f\nNumber of Particles: %d\nElement: %s\nBrush Size: %d", g.FPS, g.Particles, g.ElementSelected, g.brushSize))

	}
}
