package main

import "core:fmt"
import "core:math"
import "core:strconv"
import "core:strings"
import "core:time"
import sdl "vendor:sdl3"
import img "vendor:sdl3/image"


WINDOW_WIDTH :: 720
WINDOW_HEIGTH :: 720
WINDOW_PADDING :: 0
PATH_TO_ATLAS :: "media/atlas.png"
ART_SIZE :: 32
SDL_FLAGS :: sdl.INIT_VIDEO

state :: struct {
	window:               ^sdl.Window,
	renderer:             ^sdl.Renderer,
	atlas:                ^sdl.Texture,
	wave:                 ^wave,
	trace:                [dynamic]Vector2,
	mapStartX, mapStartY: f32,
}


drawLoop :: proc(s: ^state) {
	isRunning: bool = true
	event: sdl.Event
	for isRunning {
		for sdl.PollEvent(&event) {
			#partial switch event.type {
			case .QUIT:
				isRunning = false
				return
			}
		}

		sdl.RenderClear(s.renderer)
		sdl.RenderPresent(s.renderer)
		sdl.Delay(16)
	}
}

doInput :: proc(s: ^state) -> bool {
	event: sdl.Event
	for sdl.PollEvent(&event) {
		#partial switch event.type {
		case .QUIT:
			return true
		case:
			return false
		}
	}
	return false
}


waitForSpace :: proc(s: ^state) -> bool {
	event: sdl.Event
	for {
		for sdl.PollEvent(&event) {
			#partial switch event.type {
			case .QUIT:
				return true
			case .KEY_UP:
				if event.key.key == sdl.K_SPACE {
					return false
				}
			}
		}
		drawWave(s)
	}
}

drawTrace :: proc(s: ^state) {
	sdl.SetRenderDrawColor(s.renderer, 255, 0, 0, 255)
	signY, signX, horizontal: f32
	for v2, idx in s.trace {
		if idx + 1 >= len(s.trace) {
			break
		}
		sdl.RenderLine(
			s.renderer,
			x1 = f32(v2.x) * ART_SIZE + s.mapStartX + ART_SIZE / 2,
			y1 = f32(v2.y) * ART_SIZE + s.mapStartY + ART_SIZE / 2,
			x2 = f32(s.trace[idx + 1].x) * ART_SIZE + s.mapStartX + ART_SIZE / 2,
			y2 = f32(s.trace[idx + 1].y) * ART_SIZE + s.mapStartY + ART_SIZE / 2,
		)
		// Draw The arrow
		signY = -1
		signX = -1
		horizontal = 1
		if v2.y - s.trace[idx + 1].y > 0 {signY = 1}
		if v2.x - s.trace[idx + 1].x > 0 {signX = 1}
		if v2.y - s.trace[idx + 1].y == 0 {horizontal = -1}

		sdl.RenderLine(
			s.renderer,
			x1 = f32(s.trace[idx + 1].x) * ART_SIZE + s.mapStartX + ART_SIZE / 2,
			y1 = f32(s.trace[idx + 1].y) * ART_SIZE + s.mapStartY + ART_SIZE / 2,
			x2 = f32(s.trace[idx + 1].x) * ART_SIZE + s.mapStartX + ART_SIZE / 2 + (5 * signX),
			y2 = f32(s.trace[idx + 1].y) * ART_SIZE + s.mapStartY + ART_SIZE / 2 + (5 * signY),
		);sdl.SetRenderDrawColor(s.renderer, 0, 255, 0, 255)
		sdl.RenderLine(
			s.renderer,
			x1 = f32(s.trace[idx + 1].x) * ART_SIZE + s.mapStartX + ART_SIZE / 2,
			y1 = f32(s.trace[idx + 1].y) * ART_SIZE + s.mapStartY + ART_SIZE / 2,
			x2 = f32(s.trace[idx + 1].x) * ART_SIZE +
			s.mapStartX +
			ART_SIZE / 2 -
			(5 * signX * horizontal),
			y2 = f32(s.trace[idx + 1].y) * ART_SIZE +
			s.mapStartY +
			ART_SIZE / 2 +
			(5 * signY * horizontal),
		);sdl.SetRenderDrawColor(s.renderer, 255, 0, 0, 255)
	}
	sdl.SetRenderDrawColor(s.renderer, 255, 255, 255, 255)
}

drawWave :: proc(s: ^state) {
	if s == nil || s.wave == nil {
		fmt.panicf("drawWave got nil")
	}
	sdl.RenderClear(s.renderer)

	winW_i, winH_i: i32

	if !sdl.GetWindowSizeInPixels(s.window, &winW_i, &winH_i) {
		winH_i = WINDOW_HEIGTH
		winW_i = WINDOW_WIDTH
		fmt.println("err GetWindowSize:", sdl.GetError())
	}

	winW, winH: f32 = f32(winW_i), f32(winH_i)

	tileSize: f32 = (winH - WINDOW_PADDING) / WAVE_HEIGTH
	if winH > winW {
		tileSize = (winW - WINDOW_PADDING) / WAVE_WIDTH
	}
	tileSize = math.round_f32(tileSize)
	fmt.println(tileSize)
	drawScale := tileSize / ART_SIZE
	sdl.SetRenderScale(s.renderer, drawScale, drawScale)

	mapStartX: f32 = (winW - (tileSize * WAVE_WIDTH)) / (2 * drawScale)
	mapStartY: f32 = (winH - (tileSize * WAVE_HEIGTH)) / (2 * drawScale)
	s.mapStartX = mapStartX
	s.mapStartY = mapStartY


	src := sdl.FRect {
		x = 0,
		y = 0,
		w = ART_SIZE,
		h = ART_SIZE,
	}
	dst := sdl.FRect {
		x = mapStartX,
		y = mapStartY,
		w = ART_SIZE,
		h = ART_SIZE,
	}
	idx: int
	for row, _ in s.wave {
		for cell, _ in row {
			if getEntropy(cell) == 1 {
				idx = getIndex(cell)
				src.x = f32(idx) * ART_SIZE
				sdl.RenderTexture(s.renderer, s.atlas, &src, &dst)
			} else {
				if !sdl.SetTextureAlphaMod(s.atlas, 100) {
					fmt.printfln("err setting alpha:%v", sdl.GetError())
				}
				for v, i in cell {
					if v {
						src.x = f32(i) * ART_SIZE
						sdl.RenderTexture(s.renderer, s.atlas, &src, &dst)
					}
				}
				sdl.SetTextureAlphaMod(s.atlas, 255)
			}
			dst.x += ART_SIZE
			if dst.x >= (WAVE_WIDTH * ART_SIZE) + mapStartX {
				dst.y += ART_SIZE
				dst.x = mapStartX
			}

		}
	}
	borderRect := sdl.FRect {
		x = mapStartX - 1,
		y = mapStartY - 1,
		w = WAVE_WIDTH * ART_SIZE + 2,
		h = WAVE_HEIGTH * ART_SIZE + 2,
	}
	sdl.SetRenderDrawColor(s.renderer, 255, 0, 0, 255)
	sdl.RenderRect(s.renderer, &borderRect)
	sdl.SetRenderDrawColor(s.renderer, 255, 255, 255, 255)
	drawTrace(s)

	sdl.RenderPresent(s.renderer)
}

saveToFile :: proc(s: ^state) {
	surf := sdl.RenderReadPixels(s.renderer, nil)
	buf: [time.MIN_HMS_LEN]u8

	currTime: string = time.to_string_hms(time.now(), buf[:])
	toConcat := [?]string{"output/", currTime, ".png"}
	filePath, allocErr := strings.concatenate(toConcat[:])

	if allocErr == nil {
		file_name: cstring = strings.unsafe_string_to_cstring(filePath)
		if img.SavePNG(surf, file_name) == 0 {
			fmt.panicf("cannot SavePNG:%v", sdl.GetError())
		}
		delete(filePath)
	}
	sdl.DestroySurface(surf)

}

initSDL :: proc() {
	if !(sdl.Init(SDL_FLAGS)) {
		fmt.panicf("could not init sdl:%v", sdl.GetError())
	}

}

initState :: proc(s: ^^state) {
	s^ = new(state)
	window := sdl.CreateWindow(
		"Wave Function Collapse V3 - Odin",
		WINDOW_WIDTH,
		WINDOW_HEIGTH,
		nil,
	)
	if window == nil {
		fmt.panicf("error creating window:%v", sdl.GetError())
	}

	renderer := sdl.CreateRenderer(window, nil)
	if window == nil {
		fmt.panicf("error creating renderer:%v", sdl.GetError())
	}

	atlas := img.LoadTexture(renderer, PATH_TO_ATLAS)
	if atlas == nil {
		fmt.panicf("error loading atlas:%v", sdl.GetError())
	}

	s^.window = window
	s^.renderer = renderer
	s^.atlas = atlas
	s^.wave = nil
	s^.trace = [dynamic]Vector2{}
	sdl.SetTextureBlendMode(atlas, sdl.BLENDMODE_BLEND)
}


freeState :: proc(state: ^state) {
	delete(state.trace)
	state.trace = nil
	sdl.DestroyTexture(state.atlas)
	state.atlas = nil
	sdl.DestroyRenderer(state.renderer)
	state.renderer = nil
	sdl.DestroyWindow(state.window)
	state.window = nil
	free(state)
}

closeSDL :: proc(s: ^^state) {
	freeState(s^)
	s^ = nil
	sdl.Quit()
}
