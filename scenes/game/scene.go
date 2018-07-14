package game

import (
	"image/color"
	"math/rand"
	"time"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/Noofbiz/hypnic/collisions"
	"github.com/Noofbiz/hypnic/options"
	"github.com/Noofbiz/hypnic/systems/bullet"
	"github.com/Noofbiz/hypnic/systems/control"
	"github.com/Noofbiz/hypnic/systems/flash"
	"github.com/Noofbiz/hypnic/systems/gameover"
	"github.com/Noofbiz/hypnic/systems/gargoyle"
	"github.com/Noofbiz/hypnic/systems/gem"
	"github.com/Noofbiz/hypnic/systems/life"
	"github.com/Noofbiz/hypnic/systems/potion"
	"github.com/Noofbiz/hypnic/systems/score"
	"github.com/Noofbiz/hypnic/systems/speed"
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
	engo.Files.Load("bg.png", "mininicular.png", "vignette.png", "player.png",
		"gargoyle.png", "bullet.png", "health.png", "emptyHealth.png", "potion.png",
		"kenpixel_square.ttf", "gem.png", "bgm.mp3", "potion.wav", "gem.wav",
		"pew.wav", "hit.wav")
}

// Setup adds everything to the game to get started
func (s *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	common.SetBackground(color.Black)

	rand.Seed(time.Now().UnixNano())

	// Add Render System
	// To be added to the render system needs
	// ecs.BasicEntity
	// common.SpaceComponent
	// common.RenderComponent
	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	// add audio system
	var audioable *common.Audioable
	var notaudioable *common.NotAudioable
	w.AddSystemInterface(&common.AudioSystem{}, audioable, notaudioable)

	// add wall system
	// ecs.BasicEntity
	// common.SpaceComponent
	// wall.Component
	var wallable *wall.Able
	w.AddSystemInterface(&wall.System{}, wallable, nil)

	// add animation system
	var animationable *common.Animationable
	var notanimationable *common.NotAnimationable
	w.AddSystemInterface(&common.AnimationSystem{}, animationable, notanimationable)

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

	// add collision system
	var collisionable *common.Collisionable
	var notcollisionable *common.NotCollisionable
	w.AddSystemInterface(&common.CollisionSystem{}, collisionable, notcollisionable)

	// add bullet system
	var bulletable *bullet.Able
	w.AddSystemInterface(&bullet.System{}, bulletable, nil)

	// add health system
	w.AddSystem(&life.System{})

	// add flash system
	var flashable *flash.Able
	w.AddSystemInterface(&flash.System{}, flashable, nil)

	// add potion system
	var potionable *potion.Able
	w.AddSystemInterface(&potion.System{}, potionable, nil)

	// add score system
	w.AddSystem(&score.System{})

	// add gem system
	var gemable *gem.Able
	w.AddSystemInterface(&gem.System{}, gemable, nil)

	// add speed system
	w.AddSystem(&speed.System{})

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
	v.RenderComponent.SetZIndex(5)
	w.AddEntity(&v)

	//add player
	ps := common.NewSpritesheetWithBorderFromFile("player.png", 34, 58, 1, 1)
	p := player{BasicEntity: ecs.NewBasic()}
	p.RenderComponent.Drawable = ps.Drawable(0)
	p.RenderComponent.SetZIndex(2)
	p.SpaceComponent = common.SpaceComponent{
		Width:  p.RenderComponent.Drawable.Width(),
		Height: p.RenderComponent.Drawable.Height(),
	}
	p.SpaceComponent.SetCenter(engo.Point{
		X: engo.GameWidth() / 2,
		Y: engo.GameHeight() / 4,
	})
	p.CollisionComponent = common.CollisionComponent{
		Main:  collisions.Player,
		Group: collisions.Player,
	}
	p.AnimationComponent = common.NewAnimationComponent(ps.Drawables(), 0.05)
	p.AnimationComponent.AddAnimation(&common.Animation{
		Name:   "flash",
		Frames: []int{0, 1, 2, 3, 2, 1, 0},
	})
	w.AddEntity(&p)

	// background music
	bp, _ := common.LoadedPlayer("bgm.mp3")
	b := bgm{BasicEntity: ecs.NewBasic()}
	b.AudioComponent = common.AudioComponent{
		Player: bp,
	}
	b.AudioComponent.Player.Repeat = true
	b.AudioComponent.Player.SetVolume(options.TheOptions.BGMLevel)
	if options.TheOptions.BGM {
		b.AudioComponent.Player.Play()
	}
	w.AddEntity(&b)
}
