package main

import (
	"net/http"

	"github.com/abhay-kumar/expr-x-func"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		expression := "a + b - c"
		function := exprxfunc.Compose(expression)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<html>
				<head>
					<style>
						div {
							font-family:    Arial, Helvetica, sans-serif;
							font-size:      24px;
						}
					</style>
				</head>
				<body>
					<h1>Expression </h1><div>`+expression+`</div>
					<h1>Function </h1><div>`+function+`</div>
				</body>
			</html>
		`))
	})
	_ = r.Run()
}
