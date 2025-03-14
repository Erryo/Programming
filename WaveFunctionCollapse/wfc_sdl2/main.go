package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("WFC", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WIN_W, WIN_H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	font, err := ttf.OpenFont("media/ka1.ttf", 32)
	if err != nil {
		panic(err)
	}
	zero, err := font.RenderUTF8Solid("0", sdl.Color{255, 0, 0, 0})
	one, err := font.RenderUTF8Solid("1", sdl.Color{255, 0, 0, 0})
	two, err := font.RenderUTF8Solid("2", sdl.Color{255, 0, 0, 0})
	three, err := font.RenderUTF8Solid("3", sdl.Color{255, 0, 0, 0})
	four, err := font.RenderUTF8Solid("4", sdl.Color{255, 0, 0, 0})
	five, err := font.RenderUTF8Solid("5", sdl.Color{255, 0, 0, 0})
	six, err := font.RenderUTF8Solid("6", sdl.Color{255, 0, 0, 0})
	t0, err := renderer.CreateTextureFromSurface(zero)
	t1, err := renderer.CreateTextureFromSurface(one)
	t2, err := renderer.CreateTextureFromSurface(two)
	t3, err := renderer.CreateTextureFromSurface(three)
	t4, err := renderer.CreateTextureFromSurface(four)
	t5, err := renderer.CreateTextureFromSurface(five)
	t6, err := renderer.CreateTextureFromSurface(six)
	zero.Free()
	one.Free()
	two.Free()
	three.Free()
	four.Free()
	five.Free()
	six.Free()

	state := state{renderer: renderer, window: window}
	(&state).loadAtlas()
	state.tileAtlas.SetBlendMode(sdl.BLENDMODE_BLEND)
	wave := createWave()
	state.wave = &wave

	state.textTexture = [](*sdl.Texture){t0, t1, t2, t3, t4, t5, t6}

	arg := os.Args[1:]
	if len(arg) > 0 && arg[0] == "s" {
		side()
	}
	go generateMap(&state)

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				sdl.Quit()
				return
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_SPACE:
					}
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)

		state.drawWave()
		renderer.Present()
	}
}
