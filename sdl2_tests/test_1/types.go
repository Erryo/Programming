package main

import "github.com/veandco/go-sdl2/sdl"

// Better called level state
type gameState struct {
	window       *sdl.Window
	surface      *sdl.Surface
	gameObjects  *[]Object
	cameraTarget *Object
	nextID       uint16
}
type Object struct {
	id        uint16
	rect      sdl.Rect
	color     sdl.Color
	pixel     uint32
	direction int8
}

// Colors as pixel
var (
	MAGENTA uint32
	ORANGE  uint32
	YELLOW  uint32
	BLUE    uint32
	CYAN    uint32
	RED     uint32
)
