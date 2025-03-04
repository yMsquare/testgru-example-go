package main

import (
	"flag"
	"os"
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
)

func TestMain(t *testing.T) {
	// Save original command line args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test cases
	tests := []struct {
		name           string
		args           []string
		expectedFirst  int
		expectedSecond int
		expectedOp     string
	}{
		{
			name:           "Default values",
			args:           []string{"prog"},
			expectedFirst:  1,
			expectedSecond: 2,
			expectedOp:     "sum",
		},
		{
			name:           "Custom values",
			args:           []string{"prog", "-first", "5", "-second", "3", "-operation", "sub"},
			expectedFirst:  5,
			expectedSecond: 3,
			expectedOp:     "sub",
		},
		{
			name:           "Times operation",
			args:           []string{"prog", "-first", "4", "-second", "6", "-operation", "times"},
			expectedFirst:  4,
			expectedSecond: 6,
			expectedOp:     "times",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			flag.CommandLine = flag.NewFlagSet(tt.args[0], flag.ExitOnError)
			firstNumber = flag.Int("first", 1, "first number")
			secondNumber = flag.Int("second", 2, "second number")
			operation = flag.String("operation", "sum", "operation")

			// Set command line args
			os.Args = tt.args

			// Run main logic
			flag.Parse()
			m := math.NewMath(*firstNumber, *secondNumber, *operation)

			// Verify flag values
			if *firstNumber != tt.expectedFirst {
				t.Errorf("first number = %v, want %v", *firstNumber, tt.expectedFirst)
			}
			if *secondNumber != tt.expectedSecond {
				t.Errorf("second number = %v, want %v", *secondNumber, tt.expectedSecond)
			}
			if *operation != tt.expectedOp {
				t.Errorf("operation = %v, want %v", *operation, tt.expectedOp)
			}

			// Verify math object
			if m.A != tt.expectedFirst {
				t.Errorf("Math.A = %v, want %v", m.A, tt.expectedFirst)
			}
			if m.B != tt.expectedSecond {
				t.Errorf("Math.B = %v, want %v", m.B, tt.expectedSecond)
			}
			if m.Operator != tt.expectedOp {
				t.Errorf("Math.Operator = %v, want %v", m.Operator, tt.expectedOp)
			}
		})
	}
}
