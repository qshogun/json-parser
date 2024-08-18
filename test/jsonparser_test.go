package jsonparser_test

import (
	os "os"
	"testing"

	"github.com/qshogun/jsonparser/pkg/jsonparser"
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
			function: jsonparser.TryParse,
			file:     "testdata/step1/valid.json",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ = os.ReadFile(test.file)
			result := test.function(string(data))
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
