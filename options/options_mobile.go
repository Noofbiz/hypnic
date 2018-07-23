//+build mobilebind

package options

func LoadOptions() error {
	return nil
}

func SaveOptions() error {
	SaveNeeded = true
	return nil
}

var SaveNeeded bool
