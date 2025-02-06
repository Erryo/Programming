package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gameData struct {
	window      *sdl.Window
	surface     *sdl.Surface
	gameObjects *[]gameObject
}
type gameObject struct {
	rect      sdl.Rect
	color     sdl.Color
	pixel     uint32
	direction int8
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	gameData := gameData{window: window, surface: surface, gameObjects: &[]gameObject{}}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{120, 0, 20, 20}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

	player := gameObject{rect: rect, color: colour, pixel: pixel, direction: 10}
	*(gameData.gameObjects) = append(*gameData.gameObjects, player)

	colour2 := sdl.Color{R: 255, G: 23, B: 123, A: 255} // purple
	pixel2 := sdl.MapRGBA(surface.Format, colour2.R, colour2.G, colour2.B, colour2.A)
	rect2 := sdl.Rect{400, 300, 40, 40}

	obj := gameObject{rect: rect2, color: colour2, pixel: pixel2, direction: 0}
	*(gameData.gameObjects) = append(*gameData.gameObjects, obj)

	window.UpdateSurface()

	running := true
	frame := 0
	for running {
		frame++
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		if frame == 20 {
			frame = 0
			Update(&gameData)
			println((*gameData.gameObjects)[0].rect.X)
			drawAllGameObjects(gameData)
		}

	}
	println("Stop")
	window.UpdateSurface()
}

func drawAllGameObjects(gameData gameData) {
	gameData.surface.FillRect(nil, 0)

	for _, obj := range *gameData.gameObjects {
		gameData.surface.FillRect(&obj.rect, obj.pixel)
	}
	gameData.window.UpdateSurface()
}

func Update(gameData *gameData) {
	obj := &(*gameData.gameObjects)[0]
	rect := &obj.rect
	if rect.X+int32(obj.direction) > 700 && obj.direction > 0 {
		obj.direction *= -1
	}
	if rect.X+int32(obj.direction) < 10 && obj.direction < 0 {
		obj.direction *= -1
	}
	rect.X = rect.X + int32(obj.direction)

	// Kinda Update
}
