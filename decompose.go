package exprxfunc

import (
	"strings"
)

func Decompose(function string) string {
	// convert function to prefix expression
	prefix := funcToExp(function)
	// convert prefix to slice
	prefixSlice := strings.Fields(prefix)
	// reverse the slice
	prefixSlice = reverseSlice(prefixSlice)
	// for each slice element
	stack := NewStack()
	for _, exp := range prefixSlice {
		if isOperand(exp) {
			stack.Push(exp)
		} else if isOperator(exp) {
			a := stack.Pop().(string)
			b := stack.Pop().(string)
			infix := a + Space + exp + Space + b
			stack.Push(infix)
		}
	}
	return stack.Pop().(string)
}

func funcToExp(function string) string {
	function = strings.ReplaceAll(function, AddOpText + LeftB, AddOp + Space)
	function = strings.ReplaceAll(function, SubOpText + LeftB, SubOp + Space)
	function = strings.ReplaceAll(function, MulOpText + LeftB, MulOp + Space)
	function = strings.ReplaceAll(function, DivOpText + LeftB, DivOp + Space)
	function = strings.ReplaceAll(function, RightB + Comma, Space)
	function = strings.ReplaceAll(function, Comma, Space)
	function = strings.ReplaceAll(function, RightB, Empty)
	return function
}
