package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("lenghts are not equal\nactual: %v\nexpected:%v", len(actual), len(test.expected))
			continue
		}
		for i, actualWord := range actual {
			expectedWord := test.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal to %v", actualWord, expectedWord)
			}
		}
	}
}
