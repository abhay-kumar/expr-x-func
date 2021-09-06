package exprxfunc

func isOperator(char string) bool {
	return stringInSlice(char, []string{AddOp, SubOp, MulOp, DivOp, PowOp, LeftB, RightB})
}

func isOperand(char string) bool {
	return !stringInSlice(char, []string{AddOp, SubOp, MulOp, DivOp, PowOp, LeftB, RightB})
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func reverseSlice(s []string) []string{
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func operatorFunction(op string) string {
	switch op {
	case AddOp:
		return AddOpText
	case SubOp:
		return SubOpText
	case MulOp:
		return MulOpText
	case DivOp:
		return DivOpText
	}
	return op
}

func isOperatorFunction(op string) bool {
	return stringInSlice(op, []string{AddOpText, SubOpText, MulOpText, DivOpText})
}

func getPriority(value string) int {
	if value == SubOp || value == AddOp {
		return 1
	} else if value == MulOp || value == DivOp {
		return 2
	} else if value == PowOp {
		return 3
	}
	return 0
}
