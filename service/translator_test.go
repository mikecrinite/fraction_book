package service

import "testing"

func TestTranslateTextToDecimal(t *testing.T) {
	table := []struct {
		input string // text
		expected string // expected result
	}{
		{"Test", "0.020055069070"},
		{"zzzzzzzzzzzzzzzz", "0.076076076076076076076076076076076076076076076076"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			str := TranslateTextToDecimalString(test.input)

			if str != test.expected {
				t.Errorf("Expected=" + test.expected + ", Actual=" + str)
			}
		})
	}
}