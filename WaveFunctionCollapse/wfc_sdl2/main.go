package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Shape Generator", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	sdl.Quit()
	generateMap()

	running := false

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

		renderer.Present()
	}
}
