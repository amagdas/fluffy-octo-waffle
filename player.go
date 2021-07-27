package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

const (
	playerSpeed        = 0.05
	playerSize         = 105
	playerShotCooldown = time.Millisecond * 250
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")

	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize/2.0
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	// convert player coordinates to top left from middle of the sprite
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) < playerShotCooldown {
			return
		}
		p.shoot(p.x+25, p.y-20)
		p.shoot(p.x-25, p.y-20)
		p.lastShot = time.Now()
	}
}

func (p *player) shoot(x, y float64) {
	if bul, ok := getBulletFromPool(); ok {
		bul.active = true
		bul.x = x
		bul.y = y
		bul.angle = 270 * (math.Pi / 180)
	}
}
