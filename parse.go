package exprxfunc

import (
	"strings"
)

func infixToPostfix(infix string) string {
	var output string
	infix = LeftB + Space + infix + Space + RightB
	infixSlice := strings.Fields(infix)
	charStack := NewStack()
	for i := 0; i < len(infixSlice); i++ {
		if isOperand(infixSlice[i]) {
			output = output + infixSlice[i] + Space
		} else if infixSlice[i] == LeftB {
			charStack.Push(LeftB)
		} else if infixSlice[i] == RightB {
			for charStack.Top() != LeftB && charStack.Top() != Empty {
				output = output + charStack.Top() + Space
				charStack.Pop()
			}
			charStack.Pop()
		} else if isOperator(charStack.Top()) {
			if infixSlice[i] == PowOp {
				for getPriority(infixSlice[i]) <= getPriority(charStack.Top()) {
					output = output + charStack.Top() + Space
					charStack.Pop()
				}
			} else {
				for getPriority(infixSlice[i]) < getPriority(charStack.Top()) {
					output = output + charStack.Top() + Space
					charStack.Pop()
				}
			}
			charStack.Push(infixSlice[i])
		}
	}
	return output
}

func infixToPrefix(infix string) string {
	infixSlice := strings.Fields(infix)
	infixSlice = reverseSlice(infixSlice)
	for i := 0; i < len(infixSlice); i++ {
		if infixSlice[i] == LeftB {
			infixSlice[i] = RightB
			i++
		} else if infixSlice[i] == RightB {
			infixSlice[i] = LeftB
			i++
		}
	}
	infix = strings.Join(infixSlice, Space)
	prefix := infixToPostfix(infix)
	prefixSlice := strings.Fields(prefix)
	prefixSlice = reverseSlice(prefixSlice)
	return strings.Join(prefixSlice, Space)
}
