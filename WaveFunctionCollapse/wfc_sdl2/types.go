package main

import "github.com/veandco/go-sdl2/sdl"

type state struct {
	renderer  *sdl.Renderer
	window    *sdl.Window
	tileAtlas *sdl.Texture
	wave      *[MAP_H][MAP_W][TOTAL_TILES]bool
}

type TileModel struct {
	Index int
	Top   [TOTAL_TILES]bool
	Left  [TOTAL_TILES]bool
	Down  [TOTAL_TILES]bool
	Right [TOTAL_TILES]bool
}

const (
	WIN_W       = 1020
	WIN_H       = 900
	MAP_W       = 10
	MAP_H       = 10
	TOTAL_TILES = 5
)
