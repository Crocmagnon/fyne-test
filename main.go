package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, world")
	entry, label, errLabel := makeUI()
	w.SetContent(container.NewVBox(entry, label, errLabel))
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}

func makeUI() (*widget.Entry, *widget.Label, *widget.Label) {
	str := binding.NewString()
	in := widget.NewEntryWithData(str)
	in.Validator = validation.NewRegexp(`^\w+$`, "invalid name")

	str.Set("world")

	formatted := binding.StringToStringWithFormat(str, "Hello, %v!")
	out := widget.NewLabelWithData(formatted)

	errMsg := binding.NewString()
	in.SetOnValidationChanged(func(err error) {
		if err == nil {
			errMsg.Set("")
			return
		}

		errMsg.Set(err.Error())
	})
	errLabel := widget.NewLabelWithData(errMsg)

	return in, out, errLabel
}
