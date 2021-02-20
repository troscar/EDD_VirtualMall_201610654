package main

import (
	"./matriz"
)
func main() {
	mm := matriz.NewMatriz()
	mm.Insert("A","hogar")
	mm.Insert("C","jardin")
	mm.Insert("E","garage")
	mm.Insert("C","deporte")
	mm.Insert("F","deporte")
	mm.Insert("F","jardin")
	mm.Insert("C","garage")
	mm.Print()
}