package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/goholic/go-calculator/internal/calculator"
)

type CalculatorUI struct {
	engine *calculator.Engine
	output *widget.Label
	window fyne.Window
}

func NewCalculatorUI(w fyne.Window, e *calculator.Engine) *CalculatorUI {
	return &CalculatorUI{
		window: w,
		engine: e,
		output: widget.NewLabel("0"),
	}
}

func (ui *CalculatorUI) MakeUi() fyne.CanvasObject {
	ui.output.Alignment = fyne.TextAlignTrailing
	ui.output.TextStyle = fyne.TextStyle{Monospace: true}

	numBtn := func(text string) *widget.Button {
		return widget.NewButton(text, func() {
			ui.engine.InputNumber(text)
			ui.output.SetText(ui.engine.Display)
		})
	}

	opBtn := func(text string, op calculator.Operation) *widget.Button {
		return widget.NewButton(text, func() {
			ui.engine.InputOperator(op)
			ui.output.SetText(ui.engine.Display)
		})
	}

	calcBtn := widget.NewButton("=", func() {
		ui.engine.Calculate()
		ui.output.SetText(ui.engine.Display)
	})

	clearBtn := widget.NewButton("C", func() {
		ui.engine.Clear()
		ui.output.SetText(ui.engine.Display)
	})

	buttons := container.NewGridWithColumns(4,
		numBtn("7"), numBtn("8"), numBtn("9"), opBtn("/", calculator.OpDivide),
		numBtn("4"), numBtn("5"), numBtn("6"), opBtn("*", calculator.OpMultiply),
		numBtn("1"), numBtn("2"), numBtn("3"), opBtn("-", calculator.OpSubtract),
		clearBtn, numBtn("0"), numBtn("."), opBtn("+", calculator.OpAdd),
	)

	finalLayout := container.NewBorder(
		container.NewPadded(ui.output),
		calcBtn,
		nil, nil,
		buttons,
	)

	return finalLayout

}
