package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "What ARE you dOiNG?  ",
			expected: []string{"what", "are", "you", "doing?"},
		},
		{
			input:    "  lETS GO PiKaChu!!!",
			expected: []string{"lets", "go", "pikachu!!!"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of slices do not match")
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %v does not match expected word %v", word, expectedWord)
				return
			}
		}
	}
}
