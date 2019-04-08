//+build mobilebind

package hypnic

import (
	"github.com/EngoEngine/engo"

	"github.com/Noofbiz/hypnic/scenes/credits"
	"github.com/Noofbiz/hypnic/scenes/game"
	"github.com/Noofbiz/hypnic/scenes/mainmenu"
	"github.com/Noofbiz/hypnic/scenes/opts"
)

func Start(width, height int) {
	var gscale float32
	var cX float32
	var cY float32
	scaleX := float32(width) / 320
	scaleY := float32(height) / 480
	if scaleX > scaleY {
		gscale = scaleY
		cX = (float32(width) - 320*gscale)
	} else {
		gscale = scaleX
		cY = (float32(height) - 480*gscale)
	}
	engo.RegisterScene(&mainmenu.Scene{
		XCenter: cX,
		YCenter: cY,
	})
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
	}, &mainmenu.Scene{
		XCenter: cX,
		YCenter: cY,
	})
}
