package math_test

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
)

func TestSub(t *testing.T) {
	m := math.NewMath(10, 5, "sub")
	if m.Calculate() != 5 {
		t.Errorf("Expected 5, got %d", m.Calculate())
	}
}
