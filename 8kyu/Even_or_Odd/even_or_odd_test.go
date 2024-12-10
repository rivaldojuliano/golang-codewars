package Even_or_Odd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name     string
		request  int
		expected string
	}{
		{
			name:     "positive number",
			request:  1,
			expected: "Odd",
		},
		{
			name:     "positive number",
			request:  2,
			expected: "Even",
		},
		{
			name:     "negative number",
			request:  -1,
			expected: "Odd",
		},
		{
			name:     "negative number",
			request:  -2,
			expected: "Even",
		},
		{
			name:     "zero number",
			request:  0,
			expected: "Even",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EvenOrOdd(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
