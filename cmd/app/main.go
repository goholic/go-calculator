package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/goholic/go-calculator/internal/calculator"
	"github.com/goholic/go-calculator/ui"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GoCalc")

	engine := calculator.NewEngine()

	calcUI := ui.NewCalculatorUI(myWindow, engine)

	myWindow.SetContent(calcUI.MakeUi())
	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.ShowAndRun()
}
