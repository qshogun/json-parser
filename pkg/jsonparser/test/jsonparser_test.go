package jsonparser_test

import (
	commands "jsonparser/pkg/jsonparser"
	"os"
	"testing"
)

func TestJsonParse(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) bool
		file     string
		expected bool
	}{
		{
			name:     "IsHappyStep1Valid",
			function: commands.TryParse,
			file:     "testdata/step1/valid.json",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := os.ReadFile(test.file)
			result := test.function(string(data))
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
