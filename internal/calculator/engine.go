package calculator

import "strconv"

// Operation defines the math action
type Operation int

const (
	OpNone     Operation = iota // OP 0
	OpAdd                       // OP 1
	OpSubtract                  // OP 2
	OpMultiply                  // OP 3
	OpDivide                    // OP 4
)

// Engine holds the state of the calculator
type Engine struct {
	Value      float64   // curr num val
	Display    string    // on scrren
	currentOp  Operation // selected OP
	awaitingOp bool      // waiting for num after op?
}

// NewEngine initializes the calculator
func NewEngine() *Engine {
	return &Engine{
		Display: "0",
	}
}

// InputNumber handles digit presses (0-9) and dots
func (e *Engine) InputNumber(n string) {
	// 1. If we just hit +, -, *, /, clear the screen for the new number
	if e.awaitingOp {
		e.Display = ""
		e.awaitingOp = false
	}

	// 2. Handle the decimal point
	if n == "." {
		// If display already has a dot, ignore this input
		for _, char := range e.Display {
			if char == '.' {
				return
			}
		}
		e.Display += "."
		return
	}

	// 3. Handle Zero (Replace "0" with new number, unless it's just "0")
	if e.Display == "0" {
		e.Display = n
		return
	}

	// 4. Standard Case: Just append the number
	e.Display += n
}

// InputOperator handles +, -, *, /
func (e *Engine) InputOperator(op Operation) {
	if e.currentOp != OpNone && !e.awaitingOp {
		e.Calculate()
	}

	val, _ := strconv.ParseFloat(e.Display, 64)
	e.Value = val
	e.currentOp = op
	e.awaitingOp = true
}

// Calculate executes the pending operation
func (e *Engine) Calculate() {
	currentVal, _ := strconv.ParseFloat(e.Display, 64)

	switch e.currentOp {
	case OpAdd:
		e.Value += currentVal
	case OpSubtract:
		e.Value -= currentVal
	case OpMultiply:
		e.Value *= currentVal
	case OpDivide:
		if currentVal != 0 {
			e.Value /= currentVal
		} else {
			e.Display = "Error"
			e.currentOp = OpNone
			e.awaitingOp = true
			return
		}
	}

	e.Display = strconv.FormatFloat(e.Value, 'f', -1, 64)
	e.currentOp = OpNone
	e.awaitingOp = true
}

// clear resets the state
func (e *Engine) Clear() {
	e.Display = "0"
	e.Value = 0
	e.currentOp = OpNone
	e.awaitingOp = false
}
