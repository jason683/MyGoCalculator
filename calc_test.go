package calc

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	a, b := 6, 5
	mresult := Multiply(a, b)
	if mresult <= 0 {
		t.Error("Value is negative")
	}
	if mresult != 34 {
		t.Errorf("Value of %d is incorrect. Want 30!", mresult)
	} else {
		fmt.Println("Result is correct")
	}
}

func TestDivide(t *testing.T) {
	a, b := 6, 5
	dresult := Divide(a, b)
	if dresult <= 0 {
		t.Error("Value is negative")
	}
	if dresult != 1.2 {
		t.Errorf("Value of %.1f is incorrect. Want 1.2!", dresult)
	} else {
		fmt.Println("Result is correct")
	}
}
