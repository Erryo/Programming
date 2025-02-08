package main

func drawAllGameObjects(state gameState) {
	state.renderer.Present()
}

func prepareScene(state gameState) {
	state.renderer.SetDrawColor(255, 255, 255, 255)
	state.renderer.Clear()
}
