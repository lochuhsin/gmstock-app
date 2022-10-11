package main

import (
	"fyne.io/fyne/v2/app"
	"myapp/internal"
)

func main() {
	myApp := app.New()

	window := myApp.NewWindow("Finance Product App")

	internal.Entry(window)
}
