//+build mobilebind

package hypnic

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/scenes/credits"
	"github.com/Noofbiz/hypnic/scenes/game"
	"github.com/Noofbiz/hypnic/scenes/mainmenu"
	"github.com/Noofbiz/hypnic/scenes/opts"
)

func Start(width, height int) {
	var gscale float32
	scaleX := width / 320
	scaleY := height / 480
	if scaleX > scaleY {
		gscale = float32(scaleY)
	} else {
		gscale = float32(scaleX)
	}
	engo.RegisterScene(&mainmenu.Scene{})
	engo.RegisterScene(&game.Scene{})
	engo.RegisterScene(&credits.Scene{})
	engo.RegisterScene(&opts.Scene{})
	engo.Run(engo.RunOptions{
		Title:        "Mouse Demo",
		Width:        320,
		Height:       480,
		MobileWidth:  width,
		MobileHeight: height,
		FPSLimit:     30,
		GlobalScale: engo.Point{
			X: gscale,
			Y: gscale,
		},
	}, &mainmenu.Scene{})
}
