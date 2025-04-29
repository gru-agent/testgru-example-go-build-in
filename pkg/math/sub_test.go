package math_test

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
)

func TestSub(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},
		{0, 0, 0},
		{-1, 1, -2},
		{100, 50, 50},
		{5, 10, -5},
		{-5, -3, -2},
		{1000000, 1, 999999},
	}

	for _, tt := range tests {
		if got := math.Sub(tt.a, tt.b); got != tt.expected {
			t.Errorf("Sub(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}
