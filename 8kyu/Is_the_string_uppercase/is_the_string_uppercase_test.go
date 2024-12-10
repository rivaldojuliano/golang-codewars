package Is_the_string_uppercase

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMyString_IsUpperCase(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected bool
	}{
		{
			name:     "a",
			request:  "a",
			expected: false,
		},
		{
			name:     "WHAT DOES THE FOX SAY",
			request:  "WHAT DOES THE FOX SAY",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MyString(test.request).IsUpperCase()
			require.Equal(t, test.expected, result)
		})
	}
}
