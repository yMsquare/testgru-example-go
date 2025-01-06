package main

import (
	"flag"
	"fmt"

	"github.com/fatjyc/hello-go/pkg/math"
)

var firstNumber = flag.Int("first", 1, "first number")
var secondNumber = flag.Int("second", 2, "second number")
var operation = flag.String("operation", "sum", "operation")

func main() {
	flag.Parse()
	math := math.NewMath(*firstNumber, *secondNumber, *operation)
	fmt.Printf("Calculating %d %s %d, Result: %d\n", math.A, math.Operator, math.B, math.Calculate())
}
