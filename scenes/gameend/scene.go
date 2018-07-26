package gameend

import (
	"bytes"
	"image/color"
	"log"
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/assets"
	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/systems/gamerestart"
	"github.com/Noofbiz/hypnic/systems/menu"
	"github.com/Noofbiz/hypnic/systems/newhighscore"
)

type Scene struct {
	Score int
}

func (s *Scene) Type() string {
	return "GameEndScene"
}

func (s *Scene) Preload() {
	filelist := []string{
		"bg.png",
		"kenpixel_square.ttf",
		"scroll.png",
		"button.png",
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

	// Add Render System
	// To be added to the render system needs
	// ecs.BasicEntity
	// common.SpaceComponent
	// common.RenderComponent
	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	// add game start system
	restart := &gamerestart.System{}
	w.AddSystem(restart)

	// add menu system
	m := &menu.System{}
	w.AddSystem(m)

	// loaded Font
	fnt := &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.White,
		Size: 24,
	}
	fnt.CreatePreloaded()

	// black font
	bfnt := &common.Font{
		URL:  "kenpixel_square.ttf",
		FG:   color.Black,
		Size: 28,
	}
	bfnt.CreatePreloaded()

	//add background
	bgs, _ := common.LoadedSprite("bg.png")
	bg := sprite{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{X: options.XOffset, Y: options.YOffset},
			Width:    320,
			Height:   480,
		},
		RenderComponent: common.RenderComponent{
			Drawable: bgs,
		},
	}
	w.AddEntity(&bg)

	// Paper texture
	ps, _ := common.LoadedSprite("scroll.png")
	p := sprite{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent.Position = engo.Point{
		X: 10 + options.XOffset,
		Y: 20 + options.YOffset,
	}
	w.AddEntity(&p)

	// Score
	t := "Score: " + strconv.Itoa(s.Score)
	sc := sprite{BasicEntity: ecs.NewBasic()}
	sc.RenderComponent.Drawable = common.Text{
		Font: bfnt,
		Text: t,
	}
	sc.RenderComponent.SetZIndex(1)
	sc.SpaceComponent.Position = engo.Point{
		X: 50 + options.XOffset,
		Y: 65 + options.YOffset,
	}
	w.AddEntity(&sc)

	// Start game button
	sgs, _ := common.LoadedSprite("button.png")
	sg := button{BasicEntity: ecs.NewBasic()}
	sg.RenderComponent.Drawable = sgs
	sg.RenderComponent.SetZIndex(2)
	sg.SpaceComponent.Position = engo.Point{
		X: 50 + options.XOffset,
		Y: 110 + options.YOffset,
	}
	sg.SpaceComponent.Width = sg.RenderComponent.Drawable.Width()
	sg.SpaceComponent.Height = sg.RenderComponent.Drawable.Height()
	w.AddEntity(&sg)
	restart.Add(&sg.BasicEntity, &sg.SpaceComponent)

	// Start game text
	sgt := sprite{BasicEntity: ecs.NewBasic()}
	sgt.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Play Again",
	}
	sgt.RenderComponent.SetZIndex(3)
	sgt.SpaceComponent.Position = engo.Point{
		X: 90 + options.XOffset,
		Y: 130 + options.YOffset,
	}
	w.AddEntity(&sgt)

	// main menu button
	mm := button{BasicEntity: ecs.NewBasic()}
	mm.RenderComponent.Drawable = sgs
	mm.RenderComponent.SetZIndex(2)
	mm.SpaceComponent.Position = engo.Point{
		X: 50 + options.XOffset,
		Y: 190 + options.YOffset,
	}
	mm.SpaceComponent.Width = mm.RenderComponent.Drawable.Width()
	mm.SpaceComponent.Height = mm.RenderComponent.Drawable.Height()
	w.AddEntity(&mm)
	m.Add(&mm.BasicEntity, &mm.SpaceComponent)

	// main menu text
	mmt := sprite{BasicEntity: ecs.NewBasic()}
	mmt.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Main Menu",
	}
	mmt.RenderComponent.SetZIndex(3)
	mmt.SpaceComponent.Position = engo.Point{
		X: 90 + options.XOffset,
		Y: 210 + options.YOffset,
	}
	w.AddEntity(&mmt)

	// high score text
	if s.Score > options.TheOptions.HighScore {
		//New High Score!
		nhss := &newhighscore.System{}
		w.AddSystem(nhss)
		nhs := sprite{BasicEntity: ecs.NewBasic()}
		nhs.RenderComponent.Drawable = common.Text{
			Font: fnt,
			Text: "New High Score!",
		}
		nhs.RenderComponent.SetZIndex(2)
		nhs.SpaceComponent.Position = engo.Point{
			X: 30 + options.XOffset,
			Y: 35 + options.YOffset,
		}
		w.AddEntity(&nhs)
		nhss.Add(&nhs.BasicEntity, &nhs.RenderComponent)
		options.TheOptions.SetHighScore(s.Score)
		options.SaveOptions()
	} else {
		hst := "High Score\n" + strconv.Itoa(options.TheOptions.HighScore)
		hs := sprite{BasicEntity: ecs.NewBasic()}
		hs.RenderComponent.Drawable = common.Text{
			Font: bfnt,
			Text: hst,
		}
		hs.RenderComponent.SetZIndex(2)
		hs.SpaceComponent.Position = engo.Point{
			X: 30 + options.XOffset,
			Y: 320 + options.YOffset,
		}
		w.AddEntity(&hs)
	}
}
