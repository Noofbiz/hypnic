package mainmenu

import (
	"bytes"
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/assets"
	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/systems/creditsbtn"
	"github.com/Noofbiz/hypnic/systems/gamestart"
	"github.com/Noofbiz/hypnic/systems/optionsbtn"
)

type Scene struct{}

func (s *Scene) Type() string {
	return "MainMenuScene"
}

func (s *Scene) Preload() {
	filelist := []string{
		"menu.mp3",
		"bg.png",
		"Gaegu-Regular.ttf",
		"gargoyle.png",
		"player.png",
		"scroll.png",
		"button.png",
		"checked.png",
		"unchecked.png",
		"raise.png",
		"lower.png",
		"sound.png",
		"sfx.wav",
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
	common.SetBackground(color.Black)

	// for canvas reasizing
	options.XOffset = 3 * (320 - engo.WindowWidth()) / (8 * engo.GetGlobalScale().X)
	options.YOffset = 3 * (480 - engo.WindowHeight()) / (8 * engo.GetGlobalScale().Y)

	// Add Render System
	// To be added to the render system needs
	// ecs.BasicEntity
	// common.SpaceComponent
	// common.RenderComponent
	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	// Add Mouse System
	var mouseable *common.Mouseable
	var notmouseable *common.NotMouseable
	w.AddSystemInterface(&common.MouseSystem{}, mouseable, notmouseable)

	// add audio system
	var audioable *common.Audioable
	var notaudioable *common.NotAudioable
	w.AddSystemInterface(&common.AudioSystem{}, audioable, notaudioable)

	// add game start System
	start := &gamestart.System{}
	w.AddSystem(start)

	// add options button system
	opts := &optionsbtn.System{}
	w.AddSystem(opts)

	// add credits system
	crds := &creditsbtn.System{}
	w.AddSystem(crds)

	//add background
	bgs, _ := common.LoadedSprite("bg.png")
	bg := background{
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
	w.AddEntity(&bg)

	// background music
	bp, _ := common.LoadedPlayer("menu.mp3")
	b := bgm{BasicEntity: ecs.NewBasic()}
	b.AudioComponent = common.AudioComponent{
		Player: bp,
	}
	b.AudioComponent.Player.SetVolume(options.TheOptions.BGMLevel)
	b.AudioComponent.Player.Repeat = true
	if options.TheOptions.BGM {
		b.AudioComponent.Player.Play()
	}
	w.AddEntity(&b)

	// Title label
	tfnt := &common.Font{
		URL:  "Gaegu-Regular.ttf",
		FG:   color.White,
		Size: 32,
	}
	tfnt.CreatePreloaded()
	t := label{BasicEntity: ecs.NewBasic()}
	t.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "hypnic",
	}
	t.RenderComponent.SetZIndex(1)
	t.SpaceComponent.Position = engo.Point{
		X: 120 + options.XOffset,
		Y: 20 + options.YOffset,
	}
	w.AddEntity(&t)

	// Gargoyle
	gs, _ := common.LoadedSprite("gargoyle.png")
	g := label{BasicEntity: ecs.NewBasic()}
	g.RenderComponent.Drawable = gs
	g.RenderComponent.SetZIndex(1)
	g.SpaceComponent.Position = engo.Point{
		X: 60 + options.XOffset,
		Y: 10 + options.YOffset,
	}
	w.AddEntity(&g)

	// Mage
	ms := common.NewSpritesheetWithBorderFromFile("player.png", 34, 58, 1, 1)
	m := label{BasicEntity: ecs.NewBasic()}
	m.RenderComponent.Drawable = ms.Drawable(0)
	m.RenderComponent.SetZIndex(1)
	m.SpaceComponent.Position = engo.Point{
		X: 225 + options.XOffset,
		Y: 15 + options.YOffset,
	}
	w.AddEntity(&m)

	// Paper texture
	ps, _ := common.LoadedSprite("scroll.png")
	p := label{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent.Position = engo.Point{
		X: 10 + options.XOffset,
		Y: 80 + options.YOffset,
	}
	w.AddEntity(&p)

	// Start game button
	msfx, _ := common.LoadedPlayer("sfx.wav")
	sgs, _ := common.LoadedSprite("button.png")
	sg := button{BasicEntity: ecs.NewBasic()}
	sg.RenderComponent.Drawable = sgs
	sg.RenderComponent.SetZIndex(2)
	sg.SpaceComponent.Position = engo.Point{
		X: 20 + options.XOffset,
		Y: 100 + options.YOffset,
	}
	sg.SpaceComponent.Width = sg.RenderComponent.Drawable.Width()
	sg.SpaceComponent.Height = sg.RenderComponent.Drawable.Height()
	sg.AudioComponent.Player = msfx
	w.AddEntity(&sg)
	start.Add(&sg.BasicEntity, &sg.MouseComponent, &sg.AudioComponent)

	// Start game text
	sgt := label{BasicEntity: ecs.NewBasic()}
	sgt.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "Start",
	}
	sgt.RenderComponent.SetZIndex(3)
	sgt.SpaceComponent.Position = engo.Point{
		X: 130 + options.XOffset,
		Y: 120 + options.YOffset,
	}
	w.AddEntity(&sgt)

	// Options button
	op := button{BasicEntity: ecs.NewBasic()}
	op.RenderComponent.Drawable = sgs
	op.RenderComponent.SetZIndex(2)
	op.SpaceComponent.Position = engo.Point{
		X: 20 + options.XOffset,
		Y: 180 + options.YOffset,
	}
	op.SpaceComponent.Width = op.RenderComponent.Drawable.Width()
	op.SpaceComponent.Height = op.RenderComponent.Drawable.Height()
	op.AudioComponent.Player = msfx
	w.AddEntity(&op)
	opts.Add(&op.BasicEntity, &op.MouseComponent, &op.AudioComponent)

	// Options text
	opt := label{BasicEntity: ecs.NewBasic()}
	opt.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "Options",
	}
	opt.RenderComponent.SetZIndex(3)
	opt.SpaceComponent.Position = engo.Point{
		X: 120 + options.XOffset,
		Y: 200 + options.YOffset,
	}
	w.AddEntity(&opt)

	// Credits
	c := button{BasicEntity: ecs.NewBasic()}
	c.RenderComponent.Drawable = sgs
	c.RenderComponent.SetZIndex(2)
	c.SpaceComponent.Position = engo.Point{
		X: 20 + options.XOffset,
		Y: 380 + options.YOffset,
	}
	c.SpaceComponent.Width = c.RenderComponent.Drawable.Width()
	c.SpaceComponent.Height = c.RenderComponent.Drawable.Height()
	c.AudioComponent.Player = msfx
	w.AddEntity(&c)
	crds.Add(&c.BasicEntity, &c.MouseComponent, &c.AudioComponent)

	// Credits text
	ct := label{BasicEntity: ecs.NewBasic()}
	ct.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "Credits",
	}
	ct.RenderComponent.SetZIndex(3)
	ct.SpaceComponent.Position = engo.Point{
		X: 110 + options.XOffset,
		Y: 400 + options.YOffset,
	}
	w.AddEntity(&ct)
}
