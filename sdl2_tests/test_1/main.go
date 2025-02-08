package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(WINDOW_WIDTH, WINDOW_HEIGHT, 0)
	if err != nil {
		panic(err)
	}
	state := gameState{window: window, renderer: renderer, gameObjects: &[]Object{}}
	gameLoop(state)
}
