package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, world")
	entry, label := makeUI()
	w.SetContent(container.NewVBox(entry, label))
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}

func makeUI() (*widget.Entry, *widget.Label) {
	in := widget.NewEntry()
	out := widget.NewLabel("Hello, world!")

	in.OnChanged = func(name string) {
		if name == "" {
			name = "world"
		}
		out.SetText(fmt.Sprintf("Hello, %v!", name))
	}
	return in, out
}
