package math_test

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
)

func TestTimes(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "negative number",
			a:        -2,
			b:        3,
			expected: -6,
		},
		{
			name:     "zero",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "both negative numbers",
			a:        -2,
			b:        -3,
			expected: 6,
		},
		{
			name:     "large numbers",
			a:        1000,
			b:        1000,
			expected: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := math.Times(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Times(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
