package main

import (
	"github.com/Erryo/WFC_v2/connections"
)

const (
	W int = 60
	H int = 60
)

func main() {
	connections.Side()
	connections.DrawAllConnections()
}
