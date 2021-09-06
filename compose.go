package exprxfunc

import (
	"strings"
)

var iterator int

func Compose(expression string) string {
	prefixExpression := infixToPrefix(expression)
	expressionTree := prepareExpressionTree(prefixExpression)
	var expressionSlice []string
	expressionSlice = prepareExpression(expressionTree, expressionSlice)
	return prepareExpressionFunction(expressionSlice)
}

func prepareExpressionTree(expression string) *Tree {
	expressionSlice := strings.Fields(expression)
	expressionTree := NewTree()
	iterator = 0
	buildTree(expressionSlice, expressionTree)
	return expressionTree
}

func buildTree(op []string, t *Tree) {
	if len(op) > iterator {
		t.AddValue(op[iterator])
		iterator++
		if isOperator(op[iterator-1]) {
			if t.Left == nil {
				t.AddLeft(NewTree())
			}
			buildTree(op, t.Left)
			if t.Right == nil {
				t.AddRight(NewTree())
			}
			buildTree(op, t.Right)
		}
	}
	return
}

func prepareExpression(t *Tree, expression []string) []string {
	if t == nil {
		return expression
	}
	expression = append(expression, operatorFunction(t.Value))
	expression = prepareExpression(t.Left, expression)
	expression = prepareExpression(t.Right, expression)
	return expression
}

func prepareExpressionFunction(expressionSlice []string) string {
	var expressionFunction string
	bracketsToClose := 0
	firstParamExists := false
	for i := 0; i < len(expressionSlice); i++ {
		if isOperatorFunction(expressionSlice[i]) {
			expressionFunction = expressionFunction + expressionSlice[i] + LeftB
			bracketsToClose++
			firstParamExists = false
		} else {
			expressionFunction = expressionFunction + expressionSlice[i]
			if !firstParamExists {
				expressionFunction = expressionFunction + Comma
				firstParamExists = true
			} else {
				expressionFunction = expressionFunction + RightB
				bracketsToClose--
				if bracketsToClose > 0 {
					expressionFunction = expressionFunction + Comma
				}
			}
		}
	}
	if last := len(expressionFunction) - 1; last >= 0 && expressionFunction[last] == ',' {
		expressionFunction = expressionFunction[:last]
	}
	for i := 0; i < bracketsToClose; i++ {
		expressionFunction = expressionFunction + RightB
	}
	return expressionFunction
}
