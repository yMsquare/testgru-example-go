package main

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
	"github.com/stretchr/testify/assert"
)

func TestMathCalculate(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		operator string
		want     int
	}{
		{
			name:     "sum operation",
			a:        5,
			b:        3,
			operator: "sum",
			want:     8,
		},
		{
			name:     "sub operation",
			a:        5,
			b:        3,
			operator: "sub",
			want:     2,
		},
		{
			name:     "times operation",
			a:        5,
			b:        3,
			operator: "times",
			want:     15,
		},
		{
			name:     "invalid operation",
			a:        5,
			b:        3,
			operator: "invalid",
			want:     0,
		},
		{
			name:     "zero values",
			a:        0,
			b:        0,
			operator: "sum",
			want:     0,
		},
		{
			name:     "negative numbers",
			a:        -5,
			b:        -3,
			operator: "sum",
			want:     -8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := math.NewMath(tt.a, tt.b, tt.operator)
			got := m.Calculate()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMathNewMath(t *testing.T) {
	m := math.NewMath(1, 2, "sum")
	assert.Equal(t, 1, m.A)
	assert.Equal(t, 2, m.B)
	assert.Equal(t, "sum", m.Operator)
}

func TestMathSum(t *testing.T) {
	m := math.NewMath(5, 3, "sum")
	assert.Equal(t, 8, m.Sum())
}

func TestMathSub(t *testing.T) {
	m := math.NewMath(5, 3, "sub")
	assert.Equal(t, 2, m.Sub())
}

func TestMathTimes(t *testing.T) {
	m := math.NewMath(5, 3, "times")
	assert.Equal(t, 15, m.Times())
}
