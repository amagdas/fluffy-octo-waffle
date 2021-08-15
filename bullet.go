package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 12
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/player_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 8,
	}
	bullet.collisions = append(bullet.collisions, col)
	bullet.tag = "bullet"

	bullet.active = false
	return bullet
}

var bulletPool []*element

func initBulletPool(size int, renderer *sdl.Renderer) {
	for i := 0; i < size; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func getBulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
