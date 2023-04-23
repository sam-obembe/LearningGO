package main

import "testing"

func TestAdd(t *testing.T) {
	//arrange
	x, y := 1, 3
	expect := 4
	//act

	result := Add(x, y)

	// assert
	if expect != result {
		t.Errorf("Failed to add %v and %v. Got %v, expected %v", x, y, result, expect)
	}
}
