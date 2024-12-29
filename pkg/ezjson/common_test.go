package ezjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValueFromJSONPath(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		jsonPath  string
		expected  interface{}
		expectErr bool
	}{
		{
			name:     "Get name",
			jsonStr:  `{"name": "John", "age": 30, "address": {"city": "New York"}}`,
			jsonPath: "name",
			expected: "John",
		},
		{
			name:      "Get age",
			jsonStr:   `{"name": "John", "age": 30, "address": {"city": "New York"}}`,
			jsonPath:  "age",
			expected:  float64(30),
			expectErr: false,
		},
		{
			name:      "Get city",
			jsonStr:   `{"name": "John", "age": 30, "address": {"city": "New York"}}`,
			jsonPath:  "address.city",
			expected:  "New York",
			expectErr: false,
		},
		{
			name:      "Get non-existent country",
			jsonStr:   `{"name": "John", "age": 30, "address": {"city": "New York"}}`,
			jsonPath:  "address.country",
			expectErr: true,
		},
		{
			name:      "Get address",
			jsonStr:   `{"name": "John", "age": 30, "address": {"city": "New York"}}`,
			jsonPath:  "address",
			expected:  map[string]interface{}{"city": "New York"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := GetValueFromJSONPath(tt.jsonStr, tt.jsonPath)
			assert.Equal(t, tt.expectErr, err != nil)
			assert.Equal(t, tt.expected, val)
		})
	}
}
