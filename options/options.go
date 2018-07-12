package options

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Options struct {
	BGM, SFX           bool
	BGMLevel, SFXLevel float64
	HighScore          int
}

func (o *Options) SetHighScore(i int) {
	o.HighScore = i
}

func (o *Options) SetBGM(b bool) {
	o.BGM = b
}

func (o *Options) SetSFX(b bool) {
	o.SFX = b
}

func (o *Options) SetBGMLevel(f float64) {
	if f <= 0 {
		o.BGMLevel = 0.000001
		return
	}
	if f >= 1 {
		o.BGMLevel = 0.999999
		return
	}
	o.BGMLevel = f
}

func (o *Options) SetSFXLevel(f float64) {
	if f <= 0 {
		o.SFXLevel = 0.000001
		return
	}
	if f >= 1 {
		o.SFXLevel = 0.999999
		return
	}
	o.SFXLevel = f
}

var TheOptions Options
var location string

func LoadOptions() error {
	// for creating executables
	//exep, err := os.Executable()
	//if err != nil {
	//	return err
	//}
	//location = filepath.Join(filepath.Dir(exep), "data", "opts.json")
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	location = filepath.Join(pwd, "data", "opts.json")
	f, err := os.Open(location)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		TheOptions = Options{
			BGM:       true,
			SFX:       true,
			BGMLevel:  0.999999,
			SFXLevel:  0.999999,
			HighScore: 0,
		}
		return nil
	}
	defer f.Close()
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, &TheOptions)
	if err != nil {
		return err
	}
	return nil
}

func SaveOptions() error {
	os.Remove(location)
	f, err := os.Create(location)
	if err != nil {
		return err
	}
	defer f.Close()
	var d []byte
	d, err = json.Marshal(TheOptions)
	if err != nil {
		return err
	}
	if _, err = f.Write(d); err != nil {
		return err
	}
	return nil
}
