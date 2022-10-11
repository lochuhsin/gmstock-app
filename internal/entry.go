package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GoHomePageButton(window fyne.Window) {
	Home(window)
}

func Entry(window fyne.Window) {
	cont := container.NewVBox(
		//layout.NewBorderLayout(), // setting layouts of a page
		widget.NewLabel("This is my first app"),
		widget.NewButton("click", func() {
			GoHomePageButton(window)
		}))
	window.SetContent(cont)
	window.Resize(fyne.NewSize(500, 500))
	window.ShowAndRun()
}
