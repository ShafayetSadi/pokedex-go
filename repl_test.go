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
			input:    "  ",
			expected: []string{},
		},
		{
			input:    " Hello ",
			expected: []string{"hello"},
		},
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Gotta catch them all ",
			expected: []string{"gotta", "catch", "them", "all"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("got len:%d, but expected len:%d", len(got), len(c.expected))
		}
		for i := range got {
			word := got[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("got: %s, but expected: %s", got, expected)
			}
		}
	}

}
