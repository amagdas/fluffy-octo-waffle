package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	playerEntity, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("create the player entity", err)
		return
	}
	defer playerEntity.tex.Destroy()

	enemy, err := newBasicEnemy(renderer, screenWidth/2.0, screenHeight/2.0)
	if err != nil {
		fmt.Println("create the basic enemy", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.Clear()

		playerEntity.draw(renderer)
		playerEntity.update()

		enemy.draw(renderer)

		renderer.Present()
	}
}
