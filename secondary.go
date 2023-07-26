// Non-Reusable code to be placed in main.go

package main

import (
    //"log"
    "github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *Game{
	g := &Game{}
	g.arr1 = make([][]int, screenHeight/PixelSize)
	g.arr2 = make([][]int, screenHeight/PixelSize)
	g.pixels = make([]byte, screenWidth * screenHeight * 4)
	for i := range g.arr1 {
		g.arr1[i] = make([]int, screenWidth/PixelSize)
		g.arr2[i] = make([]int, screenWidth/PixelSize)
	}
	return g
}

func (g *Game) DrawPixels(screen *ebiten.Image){
	for row := 0; row < len(g.arr1); row++ {
		for col := 0; col < len(g.arr1[row]); col++ {
			clr := ElementMap[g.arr1[row][col]].Color
			for y := 0; y < PixelSize; y++ {
				for x := 0; x < PixelSize; x++ {
					i := ((row*PixelSize+y)*screenWidth + (col*PixelSize + x)) * 4
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

func (g *Game) UpdateCells(){
	alive := g.AliveCells_Shuffled()
	g.reset_arr2() //Clears arr2
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