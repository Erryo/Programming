package main

import "github.com/veandco/go-sdl2/sdl"

// Better called level state
type gameState struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
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

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

// Colors
var (
	MAGENTA = [3]uint8{231, 0, 106}
	ORANGE  = [3]uint8{243, 152, 1}
	YELLOW  = [3]uint8{248, 248, 69}
	BLUE    = [3]uint8{1, 104, 183}
	CYAN    = [3]uint8{50, 103, 183}
	RED     = [3]uint8{255, 0, 0}
)
