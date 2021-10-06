# expr-x-func
Compose expressions to functions and decompose functions to expressions in Go

## Why
When working with rule engines like [govaluate](https://github.com/Knetic/govaluate), it's easy to write rules which 
includes variables of primitive data types. In order to use custom data types, most rule engines expect expressions to 
be present in the functional notation. This package provides an easy way to do so.

## What
Converting expressions to functional notation is functional composition. The opposite is functional decomposition. 
Traditionally mathematical expressions are written in infix notation. It's the compilers which convert to postfix 
notation before they are executed. Functional composition is basically postfix expressions written in functional
paradigm. 

### Functional Composition
```
a + b => + a b => sum(a, b)
```

### Functional Decomposition
```
sum(a, b) => + a b => a + b
```

## How
Look at the `compose_test.go` and `decompose_test.go` for more samples

### Functional Composition

```go
package main

import (
	"fmt"
	"github.com/abhay-kumar/expr-x-func"
)

func main() {
	expression := "amount + fee - tax"
	function := exprxfunc.Compose(expression)
	fmt.Println(function) // Prints sub(add(amount,fee),tax)
}
```

### Functional Decomposition
```go
package main

import (
	"fmt"
	"github.com/abhay-kumar/expr-x-func"
)

func main() {
	function := "sub(add(amount,fee),tax)"
	expression := exprxfunc.Decompose(function)
	fmt.Println(expression) // Prints amount + fee - tax
}
```
