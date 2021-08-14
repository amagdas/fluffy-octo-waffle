package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}
	basicEnemy.position = position
	basicEnemy.rotation = 180
	basicEnemy.active = true

	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	basicEnemy.addComponent(sr)

	return basicEnemy
}
