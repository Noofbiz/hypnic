package main

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/scenes/gameend"
)

func main() {
	// uncomment this to build a portable Executable
	// also be sure to bundle the assets and data folders with it!
	//exep, _ := os.Executable()
	//root := filepath.Join(filepath.Dir(exep), "assets")
	options.LoadOptions()
	engo.Run(engo.RunOptions{
		Width:  320,
		Height: 480,
		Title:  "hypnic",
		//AssetsRoot: root,
	}, &gameend.Scene{})
}
