package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state state) drawWave() {
	winW, winH := state.window.GetSize()

	tileSize := (winW - 80) / MAP_W
	if winH < winW {
		tileSize = (winH - 80) / MAP_H
	}
	mapStartX := (winW - tileSize*MAP_W) / 2
	mapStartY := (winH - tileSize*MAP_H) / 2

	borderRect := sdl.Rect{X: mapStartX - 1, Y: mapStartY - 1, W: MAP_W*tileSize + 2, H: MAP_H*tileSize + 2}
	state.renderer.SetDrawColor(255, 0, 0, 255)
	state.renderer.DrawRect(&borderRect)
	state.renderer.SetDrawColor(255, 255, 255, 255)

	//	empty := sdl.Rect{X: 32 * 5, Y: 0, W: tileSize, H: tileSize}
	src := sdl.Rect{X: 0, Y: 0, W: 32, H: 32}
	dst := sdl.Rect{X: mapStartX, Y: mapStartY, W: tileSize, H: tileSize}

	for _, row := range state.wave {
		for _, cell := range row {
			if getEntropy(cell) == 1 {
				idx := getIndex(cell)
				src.X = int32(idx) * 32
				state.renderer.Copy(state.tileAtlas, &src, &dst)
			} else {
				if err := state.tileAtlas.SetAlphaMod(100); err != nil {
					panic(err)
				}
				for i, v := range cell {
					if v {
						src.X = int32(i) * 32
						state.renderer.Copy(state.tileAtlas, &src, &dst)
					}
				}
				state.tileAtlas.SetAlphaMod(255)
				state.renderer.Copy(state.textTexture[getEntropy(cell)], nil, &sdl.Rect{dst.X, dst.Y, tileSize / 5, tileSize / 5})
			}

			dst.X += tileSize
			if dst.X >= (MAP_W*tileSize)+mapStartX {
				dst.X = mapStartX
				dst.Y += tileSize
			}
		}
	}

	state.renderer.SetDrawColor(255, 0, 0, 255)
	for j := range MAP_W {
		i := int32(j)
		state.renderer.DrawLine(mapStartX+(i+1)*tileSize, mapStartY, mapStartX+(i+1)*tileSize, tileSize*MAP_H+mapStartY)
		state.renderer.DrawLine(mapStartX, mapStartY+(i+1)*tileSize, tileSize*MAP_W+mapStartX, mapStartY+(i+1)*tileSize)
	}
	state.renderer.SetDrawColor(0, 255, 0, 255)
	state.renderer.DrawRect(&sdl.Rect{X: (int32(state.selected[0]) * tileSize) + mapStartX, Y: (int32(state.selected[1]) * tileSize) + mapStartY, W: tileSize, H: tileSize})
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
