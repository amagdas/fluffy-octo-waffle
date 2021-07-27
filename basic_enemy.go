package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (en basicEnemy) {
	en.tex = textureFromBMP(renderer, "sprites/basic_enemy.bmp")

	en.x = x
	en.y = y
	return en
}

func (en *basicEnemy) draw(renderer *sdl.Renderer) {
	// convert player coordinates to top left from middle of the sprite
	x := en.x - basicEnemySize/2.0
	y := en.y - basicEnemySize/2.0

	renderer.CopyEx(en.tex,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize}, 180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)
}
