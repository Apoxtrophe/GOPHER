//Definitions of Elements

package main

import (
	//"log"
	//"github.com/hajimehoshi/ebiten/v2"
	"image/color"

	"golang.org/x/image/colornames"
)

type Element struct {
	Name    string
	Color   color.RGBA
	Density int
	isFluid bool
}

var ElementMap = map[int]Element{
	0: {
		Color:   colornames.Black,
		Name:    "Void",
		Density: 0,
		isFluid: true,
	},
	6: {
		Color:   colornames.Gray,
		Name:    "Carbon",
		Density: 22,
		isFluid: false,
	},
	8: {
		Color:   colornames.Aqua,
		Name:    "Oxygen",
		Density: 1,
		isFluid: true,
	},
	14: {
		Color:   colornames.Red,
		Name:    "Silicon",
		Density: 24,
		isFluid: false,
	},
	22: {
		Color:   colornames.Cornflowerblue,
		Name:    "Titanium",
		Density: 45,
		isFluid: false,
	},
	80: {
		Color:   colornames.White,
		Name:    "Mercury",
		Density: 13,
		isFluid: true,
	},
}
