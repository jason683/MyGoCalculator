package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestCalculator(t *testing.T) {
	g := Goblin(t)
	g.Describe("Calculator", func() {
		g.It("should add two numbers ", func() {
			g.Assert(Add(100, 101)).Equal(201)
			g.Assert(Add(5, -1)).Equal(4)
		})

		g.It("should subtract two numbers", func() {
			g.Assert(Subtract(1000, 997)).Equal(3)
		})

		g.It("should multiply two numbers", func() {
			g.Assert(Multiply(859, 858)).Equal(737022)
		})

		g.It("should divide two numbers", func() {
			g.Assert(Divide(29, 2)).Equal(float64(14))
		})
	})
}
