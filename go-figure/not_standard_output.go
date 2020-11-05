package main

import "github.com/common-nighthawk/go-figure"

//go get <url пакета> - установка библиотек
func main() {
	myFigure := figure.NewFigure("That's Amaizing!", "univers", true)
	myFigure.Print()
}
