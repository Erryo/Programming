package main

func drawAllGameObjects(gameData gameState) {
	gameData.surface.FillRect(nil, 0)

	for _, obj := range *gameData.gameObjects {
		gameData.surface.FillRect(&obj.rect, obj.pixel)
	}
	gameData.window.UpdateSurface()
}

func Update(gameData *gameState) {
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
