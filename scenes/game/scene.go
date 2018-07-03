package game

import (
	"image/color"
	"math/rand"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/systems/control"
	"github.com/Noofbiz/hypnic/systems/gameover"
	"github.com/Noofbiz/hypnic/systems/gargoyle"
	"github.com/Noofbiz/hypnic/systems/wall"
)

// Scene is the scene the game is played in
type Scene struct{}

// Type returns the type of scene this is
func (s *Scene) Type() string {
	return "GameScene"
}

// Preload adds in everything the game scene needs to run
func (s *Scene) Preload() {
	engo.Files.Load("bg.png", "mininicular.png", "vignette.png", "player.png", "gargoyle.png")
}

// Setup adds everything to the game to get started
func (s *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	common.SetBackground(color.White)

	rand.Seed(time.Now().UnixNano())

	// Add Render System
	// To be added to the render system needs
	// ecs.BasicEntity
	// common.SpaceComponent
	// common.RenderComponent
	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	// add wall system
	// ecs.BasicEntity
	// common.SpaceComponent
	// wall.Component
	var wallable *wall.Able
	w.AddSystemInterface(&wall.System{}, wallable, nil)

	// add control system
	// ecs.BasicEntity
	// common.SpaceComponent
	// control.Component
	var controlable *control.Able
	w.AddSystemInterface(&control.System{}, controlable, nil)

	// add gameover system
	w.AddSystem(&gameover.System{})

	// add gargoyle system
	var gargoyleable *gargoyle.Able
	w.AddSystemInterface(&gargoyle.System{}, gargoyleable, nil)

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

	//add walls
	sprites := common.NewSpritesheetFromFile("mininicular.png", 16, 16)
	walls := [][]int{
		{9, 25, 33, 8, 24, 32},
		{17, 33, 41, 40, 16, 8},
		{9, 17, 25, 8, 16, 24},
		{17, 25, 17, 8, 40, 32},
		{33, 41, 33, 32, 40, 32},
		{9, 9, 9, 8, 8, 8},
		{17, 17, 17, 16, 16, 16},
		{25, 25, 25, 24, 24, 24},
		{41, 17, 33, 24, 32, 8},
		{9, 33, 25, 40, 40, 16},
		{17, 41, 33, 8, 24, 24},
		{25, 25, 25, 16, 40, 32},
		{33, 41, 33, 32, 32, 40},
		{9, 25, 41, 40, 24, 32},
		{17, 33, 33, 24, 40, 24},
	}
	for i := 0; i < 15; i++ {
		//Left side
		l1 := wallTile{BasicEntity: ecs.NewBasic()}
		l1.RenderComponent = common.RenderComponent{
			Drawable: sprites.Drawable(walls[i][0]),
			Scale:    engo.Point{X: 2, Y: 2},
		}
		l1.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{X: 0, Y: float32(32 * i)},
			Width:    l1.Drawable.Width() * l1.Scale.X,
			Height:   l1.Drawable.Height() * l1.Scale.Y,
		}
		w.AddEntity(&l1)
		if i < 5 {
			l2 := wallTile{BasicEntity: ecs.NewBasic()}
			l2.RenderComponent = common.RenderComponent{
				Drawable: sprites.Drawable(walls[i][1]),
				Scale:    engo.Point{X: 2, Y: 2},
			}
			l2.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{X: 0, Y: float32(32*i) - engo.GameHeight()},
				Width:    l2.Drawable.Width() * l2.Scale.X,
				Height:   l2.Drawable.Height() * l2.Scale.Y,
			}
			w.AddEntity(&l2)
			l3 := wallTile{BasicEntity: ecs.NewBasic()}
			l3.RenderComponent = common.RenderComponent{
				Drawable: sprites.Drawable(walls[i][2]),
				Scale:    engo.Point{X: 2, Y: 2},
			}
			l3.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{X: 0, Y: float32(32*i) + engo.GameHeight()},
				Width:    l3.Drawable.Width() * l3.Scale.X,
				Height:   l3.Drawable.Height() * l3.Scale.X,
			}
			w.AddEntity(&l3)
		}
		// Right side
		r1 := wallTile{BasicEntity: ecs.NewBasic()}
		r1.RenderComponent = common.RenderComponent{
			Drawable: sprites.Drawable(walls[i][3]),
			Scale:    engo.Point{X: 2, Y: 2},
		}
		r1.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{
				X: engo.GameWidth() - (r1.Drawable.Width() * r1.Scale.X),
				Y: float32(32 * i),
			},
			Width:  r1.Drawable.Width() * r1.Scale.X,
			Height: r1.Drawable.Height() * r1.Scale.Y,
		}
		w.AddEntity(&r1)
		if i < 5 {
			r2 := wallTile{BasicEntity: ecs.NewBasic()}
			r2.RenderComponent = common.RenderComponent{
				Drawable: sprites.Drawable(walls[i][4]),
				Scale:    engo.Point{X: 2, Y: 2},
			}
			r2.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{
					X: engo.GameWidth() - (r2.Drawable.Width() * r2.Scale.X),
					Y: float32(32*i) - engo.GameHeight(),
				},
				Width:  r2.Drawable.Width() * r2.Scale.X,
				Height: r2.Drawable.Height() * r2.Scale.Y,
			}
			w.AddEntity(&r2)
			r3 := wallTile{BasicEntity: ecs.NewBasic()}
			r3.RenderComponent = common.RenderComponent{
				Drawable: sprites.Drawable(walls[i][5]),
				Scale:    engo.Point{X: 2, Y: 2},
			}
			r3.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{
					X: engo.GameWidth() - (r2.Drawable.Width() * r2.Scale.X),
					Y: float32(32*i) + engo.GameHeight(),
				},
				Width:  r3.Drawable.Width() * r3.Scale.X,
				Height: r3.Drawable.Height() * r3.Scale.X,
			}
			w.AddEntity(&r3)
		}
	}

	//add vignette
	vs, _ := common.LoadedSprite("vignette.png")
	v := background{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{X: 0, Y: 0},
			Width:    320,
			Height:   480,
		},
		RenderComponent: common.RenderComponent{
			Drawable: vs,
		},
	}
	w.AddEntity(&v)

	//add player
	ps, _ := common.LoadedSprite("player.png")
	p := player{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps
	p.SpaceComponent = common.SpaceComponent{
		Width:  p.RenderComponent.Drawable.Width(),
		Height: p.RenderComponent.Drawable.Height(),
	}
	p.SpaceComponent.SetCenter(engo.Point{
		X: engo.GameWidth() / 2,
		Y: engo.GameHeight() / 4,
	})
	w.AddEntity(&p)
}
