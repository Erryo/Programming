package main

import "github.com/veandco/go-sdl2/sdl"

func gameLoop(state gameState) {
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
		prepareScene(state)
		Update(state)
		drawAllGameObjects(state)

		sdl.Delay(16)
	}
}

func initObject(state *gameState, pixel uint32, x, y, w, h int32) {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	object := Object{rect: rect, pixel: pixel, id: state.nextID}
	*state.gameObjects = append(*state.gameObjects, object)
	state.nextID += 1
}

func Update(state gameState) {
	state.renderer.SetDrawColor(0, 0, 0, 255)
}
