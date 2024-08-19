// Package ch10 contains a function Add to add two integers

package ch10

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers and returns its sum.
// According to rules described in https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
