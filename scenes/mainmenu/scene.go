package mainmenu

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/systems/creditsbtn"
	"github.com/Noofbiz/hypnic/systems/gamestart"
	"github.com/Noofbiz/hypnic/systems/musicbx"
	"github.com/Noofbiz/hypnic/systems/musicdown"
	"github.com/Noofbiz/hypnic/systems/musictext"
	"github.com/Noofbiz/hypnic/systems/musicup"
	"github.com/Noofbiz/hypnic/systems/sfxbx"
	"github.com/Noofbiz/hypnic/systems/sfxdown"
	"github.com/Noofbiz/hypnic/systems/sfxtext"
	"github.com/Noofbiz/hypnic/systems/sfxup"
	"github.com/Noofbiz/hypnic/systems/soundadjust"
)

type Scene struct{}

func (s *Scene) Type() string {
	return "MainMenuScene"
}

func (s *Scene) Preload() {
	engo.Files.Load("menu.mp3", "bg.png", "Gaegu-Regular.ttf", "gargoyle.png",
		"player.png", "scroll.png", "button.png", "checked.png", "unchecked.png",
		"raise.png", "lower.png", "sound.png")
}

func (s *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	common.SetBackground(color.Black)

	// register main menu scene
	engo.RegisterScene(&Scene{})

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

	// add music down system
	mdown := &musicdown.System{}
	w.AddSystem(mdown)

	// add music up system
	mup := &musicup.System{}
	w.AddSystem(mup)

	// add music text system
	mtxt := &musictext.System{}
	w.AddSystem(mtxt)

	// add music checkbox system
	mcbx := &musicbx.System{}
	w.AddSystem(mcbx)

	// add sound adjusting systme
	sadj := &soundadjust.System{}
	w.AddSystem(sadj)

	// add credits system
	crds := &creditsbtn.System{}
	w.AddSystem(crds)

	// add sfx checkbox system
	scbx := &sfxbx.System{}
	w.AddSystem(scbx)

	// add sfx down system
	sfxd := &sfxdown.System{}
	w.AddSystem(sfxd)

	// add sfx up system
	sfxu := &sfxup.System{}
	w.AddSystem(sfxu)

	// add sfx text system
	sfxt := &sfxtext.System{}
	w.AddSystem(sfxt)

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
	b.AudioComponent.Player.Repeat = true
	b.AudioComponent.Player.Play()
	w.AddEntity(&b)
	sadj.Add(&b.BasicEntity, &b.AudioComponent)

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
	w.AddEntity(&sg)
	start.Add(&sg.BasicEntity, &sg.MouseComponent)

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
	w.AddEntity(&c)
	crds.Add(&c.BasicEntity, &c.MouseComponent)

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

	// Music label
	bfnt := &common.Font{
		URL:  "Gaegu-Regular.ttf",
		FG:   color.Black,
		Size: 32,
	}
	bfnt.CreatePreloaded()
	ml := label{BasicEntity: ecs.NewBasic()}
	ml.RenderComponent.Drawable = common.Text{
		Font: bfnt,
		Text: "Music",
	}
	ml.RenderComponent.SetZIndex(2)
	ml.SpaceComponent.Position = engo.Point{
		X: 60,
		Y: 180,
	}
	w.AddEntity(&ml)

	// Music Checkbox
	mcbs, _ := common.LoadedSprite("checked.png")
	mcb := button{BasicEntity: ecs.NewBasic()}
	mcb.RenderComponent.Drawable = mcbs
	mcb.RenderComponent.SetZIndex(3)
	mcb.SpaceComponent.Position = engo.Point{
		X: 135,
		Y: 185,
	}
	mcb.SpaceComponent.Width = mcb.RenderComponent.Drawable.Width()
	mcb.SpaceComponent.Height = mcb.RenderComponent.Drawable.Height()
	w.AddEntity(&mcb)
	mcbx.Add(&mcb.BasicEntity, &mcb.MouseComponent, &mcb.RenderComponent)

	// Music control bg
	mcbgs, _ := common.LoadedSprite("sound.png")
	mcbg := label{BasicEntity: ecs.NewBasic()}
	mcbg.RenderComponent.Drawable = mcbgs
	mcbg.RenderComponent.SetZIndex(2)
	mcbg.RenderComponent.Scale = engo.Point{
		X: 0.8,
		Y: 0.8,
	}
	mcbg.SpaceComponent.Position = engo.Point{
		X: 40,
		Y: 210,
	}
	w.AddEntity(&mcbg)

	// Music control label
	mcl := label{BasicEntity: ecs.NewBasic()}
	mcl.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "10",
	}
	mcl.RenderComponent.SetZIndex(3)
	mcl.SpaceComponent.Position = engo.Point{
		X: 150,
		Y: 220,
	}
	w.AddEntity(&mcl)
	mtxt.Add(&mcl.BasicEntity, &mcl.RenderComponent)

	// Music down
	mds, _ := common.LoadedSprite("lower.png")
	md := button{BasicEntity: ecs.NewBasic()}
	md.RenderComponent.Drawable = mds
	md.RenderComponent.SetZIndex(3)
	md.SpaceComponent.Position = engo.Point{
		X: 36,
		Y: 219,
	}
	md.SpaceComponent.Width = md.RenderComponent.Drawable.Width()
	md.SpaceComponent.Height = md.RenderComponent.Drawable.Height()
	w.AddEntity(&md)
	mdown.Add(&md.BasicEntity, &md.MouseComponent)

	// Music raise
	mrs, _ := common.LoadedSprite("raise.png")
	mr := button{BasicEntity: ecs.NewBasic()}
	mr.RenderComponent.Drawable = mrs
	mr.RenderComponent.SetZIndex(3)
	mr.SpaceComponent.Position = engo.Point{
		X: 238,
		Y: 219,
	}
	mr.SpaceComponent.Width = mr.RenderComponent.Drawable.Width()
	mr.SpaceComponent.Height = mr.RenderComponent.Drawable.Height()
	w.AddEntity(&mr)
	mup.Add(&mr.BasicEntity, &mr.MouseComponent)

	// SFX label]
	sl := label{BasicEntity: ecs.NewBasic()}
	sl.RenderComponent.Drawable = common.Text{
		Font: bfnt,
		Text: "SFX",
	}
	sl.RenderComponent.SetZIndex(2)
	sl.SpaceComponent.Position = engo.Point{
		X: 60,
		Y: 280,
	}
	w.AddEntity(&sl)

	// SFX Checkbox
	scb := button{BasicEntity: ecs.NewBasic()}
	scb.RenderComponent.Drawable = mcbs
	scb.RenderComponent.SetZIndex(3)
	scb.SpaceComponent.Position = engo.Point{
		X: 115,
		Y: 285,
	}
	scb.SpaceComponent.Width = scb.RenderComponent.Drawable.Width()
	scb.SpaceComponent.Height = scb.RenderComponent.Drawable.Height()
	w.AddEntity(&scb)
	scbx.Add(&scb.BasicEntity, &scb.MouseComponent, &scb.RenderComponent)

	// SFX control bg
	scbg := label{BasicEntity: ecs.NewBasic()}
	scbg.RenderComponent.Drawable = mcbgs
	scbg.RenderComponent.SetZIndex(2)
	scbg.RenderComponent.Scale = engo.Point{
		X: 0.8,
		Y: 0.8,
	}
	scbg.SpaceComponent.Position = engo.Point{
		X: 40,
		Y: 310,
	}
	w.AddEntity(&scbg)

	// SFX control label
	scl := label{BasicEntity: ecs.NewBasic()}
	scl.RenderComponent.Drawable = common.Text{
		Font: tfnt,
		Text: "10",
	}
	scl.RenderComponent.SetZIndex(3)
	scl.SpaceComponent.Position = engo.Point{
		X: 150,
		Y: 320,
	}
	w.AddEntity(&scl)
	sfxt.Add(&scl.BasicEntity, &scl.RenderComponent)

	// SFX down
	sd := button{BasicEntity: ecs.NewBasic()}
	sd.RenderComponent.Drawable = mds
	sd.RenderComponent.SetZIndex(3)
	sd.SpaceComponent.Position = engo.Point{
		X: 36,
		Y: 319,
	}
	sd.SpaceComponent.Width = sd.RenderComponent.Drawable.Width()
	sd.SpaceComponent.Height = sd.RenderComponent.Drawable.Height()
	w.AddEntity(&sd)
	sfxd.Add(&sd.BasicEntity, &sd.MouseComponent)

	// SFX raise
	sr := button{BasicEntity: ecs.NewBasic()}
	sr.RenderComponent.Drawable = mrs
	sr.RenderComponent.SetZIndex(3)
	sr.SpaceComponent.Position = engo.Point{
		X: 238,
		Y: 319,
	}
	sr.SpaceComponent.Width = sr.RenderComponent.Drawable.Width()
	sr.SpaceComponent.Height = sr.RenderComponent.Drawable.Height()
	w.AddEntity(&sr)
	sfxu.Add(&sr.BasicEntity, &sr.MouseComponent)
}
