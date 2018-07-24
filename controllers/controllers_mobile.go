//+build mobilebind

package controllers

import (
	"sync"
)

var (
	AccelerometerPresent bool
	lock                 sync.RWMutex
	value                float32
)

func HasKeyboard() bool {
	return false
}

func HasMouse() bool {
	return true
}

func HasAccelerometer() bool {
	return AccelerometerPresent
}

func GetAccelerometerValue() float32 {
	lock.RLock()
	defer lock.RUnlock()
	return value
}

func NewAccelerometerValue(v float32) {
	lock.Lock()
	value = (v - 0.5) * 6
	lock.Unlock()
}
