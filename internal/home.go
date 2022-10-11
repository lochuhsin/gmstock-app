package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getUniques() {

}

func Home(window fyne.Window) {

	label := widget.NewLabel("This is home page")
	form := widget.NewForm(
		widget.NewFormItem("Search for uniques", widget.NewEntry()),
	)

	cont := container.NewVBox(
		label,
		form,
	)
	window.SetContent(cont)
	window.Resize(fyne.NewSize(500, 500))
}
