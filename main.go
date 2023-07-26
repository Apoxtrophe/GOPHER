package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	pixelSize    = 10
	debugMode = true
)

type Game struct {
	pixels    []byte
	arr1      [][]int
	arr2      [][]int
	index     int
	brushSize int
	//Debug information
	FPS             int
	Particles       int
	ElementSelected string
}

func (g *Game) Update() error {
	g.Input()
	g.UpdateCells()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPixels(screen) //Draws the contents of arr1
	g.DEBUG_INFO(*screen, debugMode)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Above & Below")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
