package exprxfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompose_Sample1(t *testing.T) {
	function := "sub(add(amount,fee),tax)"
	expression := Decompose(function)
	actual := "amount + fee - tax"
	assert.Equal(t, actual, expression)
}

func TestDecompose_Sample2(t *testing.T) {
	function := "add(sub(mul(amount,fee),tax),div(bonus,dividend))"
	expression := Decompose(function)
	actual := "amount * fee - tax + bonus / dividend"
	assert.Equal(t, actual, expression)
}
