package main

import "github.com/veandco/go-sdl2/sdl"

type state struct {
	renderer  *sdl.Renderer
	window    *sdl.Window
	tileAtlas *sdl.Texture
	wave      *[MAP_H][MAP_W][TOTAL_TILES]bool
}

const (
	WIN_W       = 1920
	WIN_H       = 1080
	MAP_W       = 10
	MAP_H       = 10
	TOTAL_TILES = 5
)
