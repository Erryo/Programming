package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state state) drawWave() {
	borderRect := sdl.Rect{X: 40, Y: 40, W: MAP_W * 32, H: MAP_H * 32}
	state.renderer.DrawRect(&borderRect)
	//
	// (WIN_W - 80)/TOTAL_TILES
	empty := sdl.Rect{X: 32 * 5, Y: 0, W: 32, H: 32}
	src := sdl.Rect{X: 0, Y: 0, W: 32, H: 32}
	dst := sdl.Rect{X: 40, Y: 40, W: 32, H: 32}

	for _, row := range state.wave {
		for _, cell := range row {
			if getEntropy(cell) == 1 {
				idx := getIndex(cell)
				src.X = int32(idx) * 32
				state.renderer.Copy(state.tileAtlas, &src, &dst)
			} else {
				state.renderer.Copy(state.tileAtlas, &empty, &dst)
			}

			dst.X += 32
			if dst.X >= (MAP_W*32)+40 {
				dst.X = 40
				dst.Y += 32
			}
		}
	}
}

func printWave(state state) {
	fmt.Println("---------------------------------------------------------------------------")
	for _, row := range state.wave {
		for _, cell := range row {
			if getEntropy(cell) == 1 {
				idx := getIndex(cell)
				fmt.Printf("[%v]", strconv.Itoa(idx))
			} else {
				fmt.Print("[-]")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("---------------------------------------------------------------------------")
}

func getIndex(cell [TOTAL_TILES]bool) int {
	for j, accepted := range cell {
		if accepted {
			return j
		}
	}
	return -1
}

func (state *state) loadAtlas() {
	var err error
	if state.tileAtlas, err = img.LoadTexture(state.renderer, "media/tileAtlas.png"); err != nil {
		panic(err)
	}
}
