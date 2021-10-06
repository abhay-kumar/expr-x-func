package exprxfunc_test

import (
	"testing"

	"github.com/abhay-kumar/expr-x-func"
	"github.com/stretchr/testify/assert"
)

func TestDecompose_Sample1(t *testing.T) {
	function := "sub(add(amount,fee),tax)"
	expression := exprxfunc.Decompose(function)
	actual := "amount + fee - tax"
	assert.Equal(t, actual, expression)
}

func TestDecompose_Sample2(t *testing.T) {
	function := "add(sub(mul(amount,fee),tax),div(bonus,dividend))"
	expression := exprxfunc.Decompose(function)
	actual := "amount * fee - tax + bonus / dividend"
	assert.Equal(t, actual, expression)
}

func TestDecompose_Sample3(t *testing.T) {
	function := "amount"
	expression := exprxfunc.Decompose(function)
	actual := "amount"
	assert.Equal(t, actual, expression)
}
