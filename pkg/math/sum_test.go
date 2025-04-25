package math_test

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "negative numbers",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "zero values",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "positive and negative",
			a:        5,
			b:        -3,
			expected: 2,
		},
		{
			name:     "large numbers",
			a:        999999,
			b:        1,
			expected: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := math.Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
