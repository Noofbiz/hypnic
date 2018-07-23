//+build !mobilebind

package controllers

func HasKeyboard() bool {
	return true
}

func HasMouse() bool {
	return true
}

func HasAccelerometer() bool {
	return false
}

func GetAccelerometerValue() float32 {
	return 0
}
