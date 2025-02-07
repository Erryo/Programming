package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	state := gameState{window: window, surface: surface, gameObjects: &[]Object{}}
	surface.FillRect(nil, 0)

	MAGENTA = sdl.MapRGBA(state.surface.Format, 231, 0, 106, 255)
	ORANGE = sdl.MapRGBA(state.surface.Format, 243, 152, 1, 255)
	YELLOW = sdl.MapRGBA(state.surface.Format, 248, 248, 69, 255)
	BLUE = sdl.MapRGBA(state.surface.Format, 1, 104, 183, 255)
	CYAN = sdl.MapRGBA(state.surface.Format, 50, 103, 183, 255)
	RED = sdl.MapRGBA(state.surface.Format, 255, 0, 0, 255)

	initObject(&state, MAGENTA, 0, 0, 60, 60)
	initObject(&state, BLUE, 60, 0, 60, 60)
	initObject(&state, ORANGE, 120, 0, 60, 60)
	initObject(&state, YELLOW, 180, 0, 60, 60)
	initObject(&state, CYAN, 240, 0, 60, 60)
	initObject(&state, RED, 240, 0, 60, 60)

	fmt.Println(state.nextID, state.gameObjects)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		Update(&state)
		drawAllGameObjects(state)

		sdl.Delay(16)
	}
	window.UpdateSurface()
}
