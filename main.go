package main

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/scenes/mainmenu"
)

func main() {
	engo.Run(engo.RunOptions{
		Width:  320,
		Height: 480,
		Title:  "hypnic",
	}, &mainmenu.Scene{})
}
