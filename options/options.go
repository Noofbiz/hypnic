package options

var (
	TheOptions       Options
	XOffset, YOffset float32

	location string
)

type Options struct {
	BGM, SFX           bool
	BGMLevel, SFXLevel float64
	HighScore          int
	Controls           string
}

func (o *Options) SetControls(s string) {
	o.Controls = s
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
