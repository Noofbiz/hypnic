//+build !mobilebind

package main

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/scenes/credits"
	"github.com/Noofbiz/hypnic/scenes/game"
	"github.com/Noofbiz/hypnic/scenes/mainmenu"
	"github.com/Noofbiz/hypnic/scenes/opts"
)

func main() {
	engo.RegisterScene(&mainmenu.Scene{})
	engo.RegisterScene(&game.Scene{})
	engo.RegisterScene(&credits.Scene{})
	engo.RegisterScene(&opts.Scene{})
	options.LoadOptions()
	engo.Run(engo.RunOptions{
		Width:        320,
		Height:       480,
		FPSLimit:     30,
		NotResizable: true,
		Title:        "hypnic",
	}, &mainmenu.Scene{})
}
