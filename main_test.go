package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	if AddTwoNumbers(12, 3) != 15 {
		t.Error("Expected 12 + 3 to equal 15")
	}
}
