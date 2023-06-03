package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "Gotta catch em all!", expected: []string{"gotta", "catch", "em", "all!"}},
		{input: "Pikachu I choose you", expected: []string{"pikachu", "i", "choose", "you"}},
	}

	for _, c := range cases {
		actual := parseInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("the lengths are not equal: %v vs %v", actual, c.expected)
			continue
		}
		for i, v := range actual {
			if v != c.expected[i] {
				t.Errorf("%v does not equal %v", actual, c.expected)
				continue
			}
		}
	}
}
