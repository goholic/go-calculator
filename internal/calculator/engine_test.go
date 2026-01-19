package calculator

import "testing"

func TestAddition(t *testing.T) {
	e := NewEngine()

	e.InputNumber("1")
	e.InputNumber("0")
	e.InputOperator(OpAdd)
	e.InputNumber("5")
	e.Calculate()

	if e.Display != "15" {
		t.Errorf("Expected 15, got %s", e.Display)
	}
}

func TestChainng(t *testing.T) {
	e := NewEngine()

	e.InputNumber("2")
	e.InputOperator(OpMultiply)
	e.InputNumber("3")
	e.InputOperator(OpAdd)

	if e.Display != "6" {
		t.Errorf("Expected intermediate result 6, got %s", e.Display)
	}

	e.InputNumber("1")
	e.Calculate()

	if e.Display != "7" {
		t.Errorf("Expected final result 7, got %s", e.Display)
	}
}
