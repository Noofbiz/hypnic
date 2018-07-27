//+build !mobilebind

package options

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadOptions() error {
	// for creating executables
	// exep, err := os.Executable()
	// if err != nil {
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
			Controls:  "Touch",
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
