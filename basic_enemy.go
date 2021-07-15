package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (en basicEnemy, err error) {
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading basic enemy sprite, %v", err)
	}
	defer img.Free()

	en.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("create basic enemy texture, %v", err)
	}
	en.x = x
	en.y = y
	return en, nil
}

func (en *basicEnemy) draw(renderer *sdl.Renderer) {
	// convert player coordinates to top left from middle of the sprite
	x := en.x - basicEnemySize/2.0
	y := en.y - basicEnemySize/2.0

	renderer.CopyEx(en.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105}, 180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)
}
