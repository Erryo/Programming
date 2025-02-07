package main

import "github.com/veandco/go-sdl2/sdl"

func initObject(state *gameState, pixel uint32, x, y, w, h int32) {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	object := Object{rect: rect, pixel: pixel, id: state.nextID}
	*state.gameObjects = append(*state.gameObjects, object)
	state.nextID += 1
}
