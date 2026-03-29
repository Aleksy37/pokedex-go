package main
import "testing"



func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: " 	CharMANder 	bUlbAsauR\n 	PikAChu\n",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	{
		input: "HeLlO\nwoRLd ",
		expected: []string{"hello", "world"},
	},
	{
		input: "Hi My Name Is Ash   Ketchum",
		expected: []string{"hi", "my", "name", "is", "ash", "ketchum"},
	},
	{
		input: "",
		expected: []string{},
	},
	{
		input: " ",
		expected: []string{},
	},
}
    for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("lengths dont match: '%v' vs '%v'", actual, c.expected)
		continue
	}
	
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
}