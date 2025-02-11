package main

import (
	"flag"
	"fmt"
	"github.com/fatjyc/hello-go/tools"

	"github.com/fatjyc/hello-go/pkg/math"
)

var firstNumber = flag.Int("first", 1, "first number")
var secondNumber = flag.Int("second", 2, "second number")
var operation = flag.String("operation", "sum", "operation")

func main() {
	flag.Parse()
	m := math.NewMath(*firstNumber, *secondNumber, *operation)
	c := m.Calculate()
	fmt.Printf("Calculating %d %s %d, Result: %d\n", m.A, m.Operator, m.B, c)

	fmt.Println("IsOdd:", tools.IsOdd(c))
	fmt.Println("IsEven:", tools.IsEven(c))
	fmt.Println("IsPrime:", tools.IsPrime(c))
}
