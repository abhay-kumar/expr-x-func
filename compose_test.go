package exprxfunc_test

import (
	"testing"

	"github.com/abhay-kumar/expr-x-func"
	"github.com/stretchr/testify/assert"
)

func TestCompose_Sample1(t *testing.T) {
	expression := "amount + fee - tax"
	function := exprxfunc.Compose(expression)
	actual := "sub(add(amount,fee),tax)"
	assert.Equal(t, actual, function)
}

func TestCompose_Sample2(t *testing.T) {
	expression := "amount * fee - tax + bonus / dividend"
	function := exprxfunc.Compose(expression)
	actual := "add(sub(mul(amount,fee),tax),div(bonus,dividend))"
	assert.Equal(t, actual, function)
}
