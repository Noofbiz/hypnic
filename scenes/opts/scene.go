package opts

import (
	"image/color"
	"math"
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/systems/backbtn"
	"github.com/Noofbiz/hypnic/systems/musicbx"
	"github.com/Noofbiz/hypnic/systems/musicdown"
	"github.com/Noofbiz/hypnic/systems/musictext"
	"github.com/Noofbiz/hypnic/systems/musicup"
	"github.com/Noofbiz/hypnic/systems/sfxadjust"
	"github.com/Noofbiz/hypnic/systems/sfxbx"
	"github.com/Noofbiz/hypnic/systems/sfxdown"
	"github.com/Noofbiz/hypnic/systems/sfxtext"
	"github.com/Noofbiz/hypnic/systems/sfxup"
	"github.com/Noofbiz/hypnic/systems/soundadjust"
)

type Scene struct{}

func (s *Scene) Type() string {
	return "OptionsScene"
}

func (s *Scene) Preload() {
	engo.Files.Load("menu.mp3", "bg.png", "Gaegu-Regular.ttf",
		"scroll.png", "button.png", "checked.png", "unchecked.png",
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

	// add sfx adjusting system
	sfxadj := &sfxadjust.System{}
	w.AddSystem(sfxadj)

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

	// add back button system
	sbbtn := &backbtn.System{}
	w.AddSystem(sbbtn)

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
		Text: "Options",
	}
	t.RenderComponent.SetZIndex(1)
	t.SpaceComponent.Position = engo.Point{
		X: 120,
		Y: 20,
	}
	w.AddEntity(&t)

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

	// sfx player
	msfx, _ := common.LoadedPlayer("sfx.wav")
	// button background texture
	sgs, _ := common.LoadedSprite("button.png")

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
	mcbsu, _ := common.LoadedSprite("unchecked.png")
	mcb := button{BasicEntity: ecs.NewBasic()}
	if options.TheOptions.BGM {
		mcb.RenderComponent.Drawable = mcbs
	} else {
		mcb.RenderComponent.Drawable = mcbsu
	}
	mcb.RenderComponent.SetZIndex(3)
	mcb.SpaceComponent.Position = engo.Point{
		X: 135,
		Y: 185,
	}
	mcb.SpaceComponent.Width = mcb.RenderComponent.Drawable.Width()
	mcb.SpaceComponent.Height = mcb.RenderComponent.Drawable.Height()
	mcb.AudioComponent.Player = msfx
	w.AddEntity(&mcb)
	mcbx.Add(&mcb.BasicEntity, &mcb.MouseComponent, &mcb.RenderComponent, &mcb.AudioComponent)

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
	mci := int(math.Round(options.TheOptions.BGMLevel * 10))
	mcl.RenderComponent.SetZIndex(3)
	mcl.SpaceComponent.Position = engo.Point{
		X: 150,
		Y: 220,
	}
	if mci == 2 {
		mcl.RenderComponent.Drawable = tfnt.Render("2")
		mcl.SpaceComponent.Position.Y += 5
	} else {
		mclt := strconv.Itoa(mci)
		mcl.RenderComponent.Drawable = common.Text{
			Font: tfnt,
			Text: mclt,
		}
	}
	w.AddEntity(&mcl)
	mtxt.Add(&mcl.BasicEntity, &mcl.SpaceComponent, &mcl.RenderComponent)

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
	md.AudioComponent.Player = msfx
	w.AddEntity(&md)
	mdown.Add(&md.BasicEntity, &md.MouseComponent, &md.AudioComponent)

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
	mr.AudioComponent.Player = msfx
	w.AddEntity(&mr)
	mup.Add(&mr.BasicEntity, &mr.MouseComponent, &mr.AudioComponent)

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
	if options.TheOptions.SFX {
		scb.RenderComponent.Drawable = mcbs
	} else {
		scb.RenderComponent.Drawable = mcbsu
	}
	scb.RenderComponent.SetZIndex(3)
	scb.SpaceComponent.Position = engo.Point{
		X: 115,
		Y: 285,
	}
	scb.SpaceComponent.Width = scb.RenderComponent.Drawable.Width()
	scb.SpaceComponent.Height = scb.RenderComponent.Drawable.Height()
	scb.AudioComponent.Player = msfx
	w.AddEntity(&scb)
	scbx.Add(&scb.BasicEntity, &scb.MouseComponent, &scb.RenderComponent, &scb.AudioComponent)

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
	scli := int(math.Round(options.TheOptions.SFXLevel * 10))
	scl.RenderComponent.SetZIndex(3)
	scl.SpaceComponent.Position = engo.Point{
		X: 150,
		Y: 320,
	}
	if scli == 2 {
		scl.RenderComponent.Drawable = tfnt.Render("2")
		scl.SpaceComponent.Position.Y += 5
	} else {
		sclt := strconv.Itoa(scli)
		scl.RenderComponent.Drawable = common.Text{
			Font: tfnt,
			Text: sclt,
		}
	}
	w.AddEntity(&scl)
	sfxt.Add(&scl.BasicEntity, &scl.SpaceComponent, &scl.RenderComponent)

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
	sd.AudioComponent.Player = msfx
	w.AddEntity(&sd)
	sfxd.Add(&sd.BasicEntity, &sd.MouseComponent, &sd.AudioComponent)

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
	sr.AudioComponent.Player = msfx
	w.AddEntity(&sr)
	sfxu.Add(&sr.BasicEntity, &sr.MouseComponent, &sr.AudioComponent)

	// backbutton
	bb := button{BasicEntity: ecs.NewBasic()}
	bb.RenderComponent.Drawable = sgs
	bb.RenderComponent.SetZIndex(2)
	bb.SpaceComponent.Position = engo.Point{
		X: 20,
		Y: 100,
	}
	bb.SpaceComponent.Width = bb.RenderComponent.Drawable.Width()
	bb.SpaceComponent.Height = bb.RenderComponent.Drawable.Height()
	bb.AudioComponent.Player = msfx
	w.AddEntity(&bb)
	sbbtn.Add(&bb.BasicEntity, &bb.MouseComponent, &bb.AudioComponent)
	sfxadj.Add(&bb.BasicEntity, &bb.AudioComponent)

	// back button text
	bbt := label{BasicEntity: ecs.NewBasic()}
	bbt.RenderComponent.Drawable = tfnt.Render("Back")
	bbt.RenderComponent.SetZIndex(3)
	bbt.SpaceComponent.Position = engo.Point{
		X: 130,
		Y: 125,
	}
	w.AddEntity(&bbt)
}
