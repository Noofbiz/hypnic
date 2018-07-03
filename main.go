package main

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/scenes/game"
)

func main() {
	engo.Run(engo.RunOptions{
		Width:  320,
		Height: 480,
		Title:  "Hypnic",
	}, &game.Scene{})
}
