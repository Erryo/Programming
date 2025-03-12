package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func side() {
	shapes := getTileShape()
	//!!!FUZZY

	models := initModels(len(shapes))
	// It is not commutative AB != BA
	// Check A against B
	for i := range shapes {
		for j := i; j < len(shapes); j++ {
			connect(i, j, shapes, &models)
		}
	}
	setTileModel(models)
}

func initModels(noTiles int) []TileModel {
	models := make([]TileModel, noTiles)
	model := TileModel{}
	for i := range models {
		model.Index = i
		model.Right = [5]bool{}
		model.Left = [5]bool{}
		model.Down = [5]bool{}
		model.Top = [5]bool{}
		models[i] = model
	}
	return models
}

func connect(idxA, idxB int, shapes [][9]bool, models *[]TileModel) {
	// 0-2 T        |012|
	// 3-5 M        |345|
	// 6-8 D        |678|
	// 0,0+3,0+6 L
	// 2,2+3,2+6 R

	if idxB == 4 {
		fmt.Println("Hit", shapes[idxA], shapes[idxB])
	}
	a := shapes[idxA]
	b := shapes[idxB]
	modelA := &(*models)[idxA]
	modelB := &(*models)[idxB]
	modelA.Index = idxA
	modelB.Index = idxB

	// Check A's Top with B's Bottom
	valid := true
	for i := 0; i <= 2; i++ {
		if a[i] != b[6+i] {
			valid = false
			break
		}
	}
	if valid {
		modelA.Top[idxB] = true
		modelB.Down[idxA] = true
	}
	valid = true

	// Check A's Bottom with B's Top
	for i := 0; i <= 2; i++ {
		if a[6+i] != b[i] {
			valid = false
			break
		}
	}
	if valid {
		modelA.Down[idxB] = true
		modelB.Top[idxA] = true
	}
	valid = true

	// Check A'- Left wih B's right
	for i := 0; i <= 6; i += 3 {
		if a[i] != b[i+2] {
			valid = false
			break
		}
	}
	if valid {
		modelA.Left[idxB] = true
		modelB.Right[idxA] = true
	}
	valid = true

	// Check A'- right wih B's left
	for i := 0; i <= 6; i += 3 {
		if a[i+2] != b[i] {
			valid = false
			break
		}
	}
	if valid {
		if idxB == 4 {
			fmt.Println("Right")
		}
		modelA.Right[idxB] = true
		modelB.Left[idxA] = true
	}
	valid = true
}

func getTileModel() []TileModel {
	models := []TileModel{}
	data, err := os.ReadFile("media/tileModel.json")
	if err != nil {
		fmt.Println("Couldnt ReadFile tileModel::", err)
		if os.IsNotExist(err) {
			fmt.Println("creating file")
			createJsonFile("media/tileModel.json")
		}
		return models
	}
	err = json.Unmarshal(data, &models)
	if err != nil {
		fmt.Println("Could not Unmarshal tileModel data::", err)
	}

	return models
}

func getTileShape() [][9]bool {
	shapes := [][9]bool{}
	data, err := os.ReadFile("media/tileShape.json")
	if err != nil {
		fmt.Println("Couldnt ReadFile tileShape::", err)
		if os.IsNotExist(err) {
			fmt.Println("creating file")
			createJsonFile("media/tileShape.json")
		}
		return shapes
	}
	err = json.Unmarshal(data, &shapes)
	if err != nil {
		fmt.Println("Could not Unmarshal tileShape data::", err)
	}

	return shapes
}

func setTileModel(models []TileModel) {
	file, err := os.OpenFile("media/tileModel.json", os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Print("error opening file: tileModel.json", err)
		return
	}
	defer file.Close()

	jsonData := convertModelToJson(models)
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("error writing tileModel.(json) to file", err)
	}
}

func setTileShape(shapes [][9]bool) {
	file, err := os.OpenFile("media/tileShape.json", os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Print("error opening file: tileShape.json", err)
		return
	}
	defer file.Close()

	jsonData := convertShapeToJson(shapes)
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("error writing tileShape.(json) to file", err)
	}
}

func convertShapeToJson(shapes [][9]bool) []byte {
	data, err := json.MarshalIndent(shapes, "", "   ")
	if err != nil {
		fmt.Println("error marshaling tileShape to Json::", err)
	}
	return data
}

func convertModelToJson(models []TileModel) []byte {
	data, err := json.MarshalIndent(models, "", "   ")
	if err != nil {
		fmt.Println("error marshaling tileModel to Json::", err)
	}
	return data
}

func createJsonFile(filepath string) {
	_, err := os.Create(filepath)
	if err != nil {
		fmt.Println("could not createJsonFile: ", filepath, "::", err)
	}
}
