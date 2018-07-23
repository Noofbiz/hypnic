//+build mobilebind

package androidglue

import (
	"encoding/json"

	"github.com/Noofbiz/hypnic"
	"github.com/Noofbiz/hypnic/controllers"
	"github.com/Noofbiz/hypnic/options"

	"engo.io/engo"
)

var running bool

func Start(width, height int) {
	running = true
	hypnic.Start(width, height)
}

func Update() {
	engo.RunIteration()
}

func IsRunning() bool {
	return running
}

func Touch(x, y, id, action int) {
	engo.TouchEvent(x, y, id, action)
}

func Stop() {
	running = false
	engo.MobileStop()
}

func LoadFromJava(b []byte) {
	if len(b) == 0 {
		options.TheOptions = options.Options{
			BGM:       true,
			SFX:       true,
			BGMLevel:  0.999999,
			SFXLevel:  0.999999,
			HighScore: 0,
			Controls:  "Touch",
		}
		return
	}
	json.Unmarshal(b, &options.TheOptions)
}

func NeedsSaving() bool {
	if options.SaveNeeded {
		options.SaveNeeded = false
		return true
	}
	return false
}

func TheOptionsBytes() []byte {
	var d []byte
	d, _ = json.Marshal(options.TheOptions)
	return d
}

func AccelerometerPresent(b bool) {
	controllers.AccelerometerPresent = b
}

func AccelerometerValue(v float32) {
	controllers.NewAccelerometerValue(v)
}
