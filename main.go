package main

import (
	"engo.io/engo"

	"github.com/Noofbiz/hypnic/scenes/mainmenu"
)

func main() {
	// uncomment this to build a portable Executable
	// also be sure to bundle the assets folder with it!
	//exep, _ := os.Executable()
	//root := filepath.Join(filepath.Dir(exep), "assets")
	engo.Run(engo.RunOptions{
		Width:  320,
		Height: 480,
		Title:  "hypnic",
		//AssetsRoot: root,
	}, &mainmenu.Scene{})
}
