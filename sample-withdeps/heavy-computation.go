package sample_withdeps

import "github.com/common-nighthawk/go-figure"

func Print(text string) {
	myFigure := figure.NewFigure(text, "", true)
	myFigure.Print()
}

func PrintHello() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}
