package mainmenu

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

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
	engo.Files.Load("menu.mp3", "bg.png", "Gaegu-Regular.ttf", "gargoyle.png",
		"player.png", "scroll.png", "button.png", "checked.png", "unchecked.png",
		"raise.png", "lower.png", "sound.png", "sfx.wav")
}

func (s *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	common.SetBackground(color.Black)

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
			Position: engo.Point{X: 0, Y: 0},
			Width:    320,
			Height:   480,
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
		X: 120,
		Y: 20,
	}
	w.AddEntity(&t)

	// Gargoyle
	gs, _ := common.LoadedSprite("gargoyle.png")
	g := label{BasicEntity: ecs.NewBasic()}
	g.RenderComponent.Drawable = gs
	g.RenderComponent.SetZIndex(1)
	g.SpaceComponent.Position = engo.Point{
		X: 60,
		Y: 10,
	}
	w.AddEntity(&g)

	// Mage
	ms := common.NewSpritesheetWithBorderFromFile("player.png", 34, 58, 1, 1)
	m := label{BasicEntity: ecs.NewBasic()}
	m.RenderComponent.Drawable = ms.Drawable(0)
	m.RenderComponent.SetZIndex(1)
	m.SpaceComponent.Position = engo.Point{
		X: 225,
		Y: 15,
	}
	w.AddEntity(&m)

	// Paper texture
	ps, _ := common.LoadedSprite("scroll.png")
	p := label{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent.Position = engo.Point{
		X: 10,
		Y: 80,
	}
	w.AddEntity(&p)

	// Start game button
	msfx, _ := common.LoadedPlayer("sfx.wav")
	sgs, _ := common.LoadedSprite("button.png")
	sg := button{BasicEntity: ecs.NewBasic()}
	sg.RenderComponent.Drawable = sgs
	sg.RenderComponent.SetZIndex(2)
	sg.SpaceComponent.Position = engo.Point{
		X: 20,
		Y: 100,
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
		X: 130,
		Y: 120,
	}
	w.AddEntity(&sgt)

	// Options button
	op := button{BasicEntity: ecs.NewBasic()}
	op.RenderComponent.Drawable = sgs
	op.RenderComponent.SetZIndex(2)
	op.SpaceComponent.Position = engo.Point{
		X: 20,
		Y: 180,
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
		X: 120,
		Y: 200,
	}
	w.AddEntity(&opt)

	// Credits
	c := button{BasicEntity: ecs.NewBasic()}
	c.RenderComponent.Drawable = sgs
	c.RenderComponent.SetZIndex(2)
	c.SpaceComponent.Position = engo.Point{
		X: 20,
		Y: 380,
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
		X: 110,
		Y: 400,
	}
	w.AddEntity(&ct)
}
