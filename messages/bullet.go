package messages

import "engo.io/engo"

// CreateBulletType is the type for the bullet message
const CreateBulletType string = "CreateBulletMessage"

// CreateBullet is a message that creates a bullet from the given location and
// fires it at the given angle in degrees (theta = 0 is vertical and increases
// counter-clockwise.)
type CreateBullet struct {
	Position engo.Point
	Angle    float32
}

// Type implements the engo.Message interface
func (CreateBullet) Type() string {
	return CreateBulletType
}
