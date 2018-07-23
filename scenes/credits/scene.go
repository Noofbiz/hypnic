package credits

import (
	"bytes"
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/assets"
	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/systems/creditroll"
	"github.com/Noofbiz/hypnic/systems/endcredits"
)

type Scene struct {
	fnt, dfnt *common.Font
	curPos    float32
	w         *ecs.World
}

func (s *Scene) Type() string {
	return "CreditsScene"
}

func (s *Scene) Preload() {
	filelist := []string{
		"bg.png",
		"bgm.mp3",
		"Gaegu-Regular.ttf",
		"kenpixel_square.ttf",
	}
	for _, url := range filelist {
		d, err := assets.Asset(url)
		if err != nil {
			log.Println("Couldn't load " + url)
		}
		engo.Files.LoadReaderData(url, bytes.NewReader(d))
	}
}

func (s *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	s.w = w
	common.SetBackground(color.Black)

	// Add Render System
	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	s.w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	// add audio system
	var audioable *common.Audioable
	var notaudioable *common.NotAudioable
	s.w.AddSystemInterface(&common.AudioSystem{}, audioable, notaudioable)

	// add credit roll system
	var creditrollable *creditroll.Able
	s.w.AddSystemInterface(&creditroll.System{}, creditrollable, nil)

	// add end credit system
	s.w.AddSystem(&endcredits.System{})

	//add background
	bgs, _ := common.LoadedSprite("bg.png")
	bg := sprite{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{
				X: 0 + options.XOffset,
				Y: 0 + options.YOffset,
			},
			Width:  320,
			Height: 480,
		},
		RenderComponent: common.RenderComponent{
			Drawable: bgs,
		},
	}
	s.w.AddEntity(&bg)

	// background music
	bgmp, _ := common.LoadedPlayer("bgm.mp3")
	bgmp.SetVolume(options.TheOptions.BGMLevel)
	b := bgm{BasicEntity: ecs.NewBasic()}
	b.AudioComponent = common.AudioComponent{
		Player: bgmp,
	}
	b.AudioComponent.Player.Repeat = true
	if options.TheOptions.BGM {
		b.AudioComponent.Player.Play()
	}
	s.w.AddEntity(&b)

	s.curPos = 320 + options.YOffset

	//up font
	s.fnt = &common.Font{
		URL:  "Gaegu-Regular.ttf",
		FG:   color.White,
		Size: 32,
	}
	s.fnt.CreatePreloaded()

	//down font
	s.dfnt = &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.White,
		Size: 24,
	}
	s.dfnt.CreatePreloaded()

	//Title
	s.createLines("hypnic")

	//me
	s.createLines("A Game By", "Jerry Caligiure", "noofbiz.github.io")

	// sprites
	s.createLines("sprites")
	// background
	s.createLines("Background", "Robin Caligiure")

	// bullet
	s.createLines("Statue Bullets", "Kenney", "kenney.nl")

	// main menu buttons
	s.createLines("Main Menu Buttons", "Under the moon", "opengameart.org/users/under-the-moon")

	// health bars
	s.createLines("Health Bars", "Scrittl", "opengameart.org/users/scrittl")

	// mage player
	s.createLines("Player Mage", "Sollision", "opengameart.org/user/30796",
		"Jordan Irwin", "opengameart.org/users/antumdeluge")

	// mage statue
	s.createLines("Mage Statue", "Johann C", "opengameart.org/users/johann-c")

	// gem
	s.createLines("Gem", "Code Inferno Games", "codeinferno.com")

	// walls
	s.createLines("Walls", "Blarget2", "opengameart.org/users/blarget2")

	// potion
	s.createLines("Potion", "Bonsaiheldin", "bonsaiheld.org")

	// scroll background
	s.createLines("Paper", "darkwood67", "deviantart.com/darkwood67/")

	// sound
	s.createLines("sounds")
	// menu bgm
	s.createLines("Menu BGM", "HorrorPen", "opengameart.org/users/horrorpen")

	// game bgm
	s.createLines("Game BGM", "xXUnderTowerXx", "opengameart.org/users/xxundertowerxx")

	// potion pickup
	s.createLines("Potion Sound", "Bart Kelsey", "opengameart.org/users/bart")

	// hit
	s.createLines("On Hit", "wobbleboxx", "wobbleboxx.com")

	// gem
	s.createLines("Gem Sound", "wobbleboxx", "wobbleboxx.com")

	// statue fire
	s.createLines("Statue Fire", "dklon", "opengameart.org/users/dklon")

	// button press
	s.createLines("Button Press", "copyc4t", "opengameart.org/users/copyc4t")

	// font
	s.createLines("fonts")
	// kenpixel
	s.createLines("Square Pixel", "Kenney", "kenney.nl")

	// Gaegu
	s.createLines("Gaegu", "JIKJI SOFT", "fonts.google.com/specimen/Gaegu")
}

func (s *Scene) createLines(lines ...string) {
	for i, line := range lines {
		l := mobilesprite{BasicEntity: ecs.NewBasic()}
		if i%2 == 1 {
			l.RenderComponent.Drawable = s.dfnt.Render(line)
			l.RenderComponent.Scale = engo.Point{
				X: 1,
				Y: 1,
			}
		} else if i == 0 {
			l.RenderComponent.Drawable = s.fnt.Render(line)
			l.RenderComponent.Scale = engo.Point{
				X: 1,
				Y: 1,
			}
		} else if i%2 == 0 {
			l.RenderComponent.Drawable = s.fnt.Render(line)
			l.RenderComponent.Scale = engo.Point{
				X: 0.5,
				Y: 0.5,
			}
		}
		l.RenderComponent.SetZIndex(2)
		l.SpaceComponent.Width = l.RenderComponent.Drawable.Width() * l.RenderComponent.Scale.X
		l.SpaceComponent.Height = l.RenderComponent.Drawable.Height() * l.RenderComponent.Scale.Y
		l.SpaceComponent.Position = engo.Point{
			X: (320-l.SpaceComponent.Width)/2 + ((320 - engo.WindowWidth()) / (2 * engo.GetGlobalScale().X)),
			Y: s.curPos,
		}
		s.w.AddEntity(&l)
		s.curPos += 28
	}
	s.curPos += 30
}
