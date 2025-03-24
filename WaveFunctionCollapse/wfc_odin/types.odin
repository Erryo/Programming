package main

WAVE_WIDTH :: 30
WAVE_HEIGTH :: 30
TOTAL_TILES :: 5
PATH_TO_MODELS :: "media/models.json"
PATH_TO_SHAPES :: "media/simpleTile.json"
Debug: bool = false

tile :: [TOTAL_TILES]bool
wave :: [WAVE_HEIGTH][WAVE_WIDTH]tile
Vector2 :: [2]int

tileModel :: struct {
	index: int `json:"Index"`,
	top:   tile `json:"Top"`,
	left:  tile `json:"Left"`,
	down:  tile `json:"Down"`,
	right: tile `json:"Right"`,
}
