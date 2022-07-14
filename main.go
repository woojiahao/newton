package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	titleText := widget.NewLabel("Newton")
	subTitleText := widget.NewLabel("Newton's Method of Approximation Visualizer")
	titleContainer := container.NewVBox(titleText, subTitleText)

	inputText := widget.NewLabel("Enter the function to visualize")
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Function...")
	startingPointText := widget.NewLabel("Enter the starting x-coordinate")
	startingPointEntry := widget.NewEntry()
	startingPointEntry.SetPlaceHolder("x-coordinate...")
	startVisualizationButton := widget.NewButton("Start", func() {
	})
	stopVisualizationButton := widget.NewButton("Stop", func() {

	})
	inputContainer := container.NewVBox(
		container.NewHBox(inputText, inputEntry),
		container.NewHBox(startingPointText, startingPointEntry),
		container.NewHBox(startVisualizationButton, stopVisualizationButton),
	)

	visualization := canvas.NewImageFromResource(theme.FyneLogo())
	visualization.FillMode = canvas.ImageFillOriginal

	w.SetContent(container.NewBorder(titleContainer, nil, nil, inputContainer, visualization))

	w.Show()

	a.Run()
}
