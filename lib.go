//Highly reusable code to be used in "secondary" functions

package main


import (
    //"log"
    //"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

//Used in UpateCells in secondary
func (g *Game) AliveCells_Shuffled()([][2]int){
	aliveCells := make([][2]int, 0)
	for row := 0; row < len(g.arr1); row++ {
		for col := 0; col < len(g.arr1[row]); col++ {
			if g.arr1[row][col] != 0 {
				aliveCells = append(aliveCells, [2]int{row, col})
			}
		}
	}
	rand.Shuffle(len(aliveCells), func(i, j int) {
		aliveCells[i], aliveCells[j] = aliveCells[j], aliveCells[i]
	})
	return aliveCells
}
//Used in UpateCells in secondary
func (g *Game)reset_arr2(){
	for row := range g.arr1 {
		for col := range g.arr1[row] {
			g.arr2[row][col] = 0
		}
	}
}

func clamp(val, min, max int) int {
    if val < min {
        return min
    } else if val > max {
        return max
    } else {
        return val
    }
}