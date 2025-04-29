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
			a:        5,
			b:        3,
			expected: 8,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -4,
			expected: -6,
		},
		{
			name:     "zero values",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "positive and negative",
			a:        10,
			b:        -5,
			expected: 5,
		},
		{
			name:     "max int",
			a:        2147483647,
			b:        1,
			expected: 2147483648,
		},
		{
			name:     "min int",
			a:        -2147483648,
			b:        -1,
			expected: -2147483649,
		},
		{
			name:     "large numbers",
			a:        999999999,
			b:        999999999,
			expected: 1999999998,
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
