package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var testX, testY, testResult int

// Save the testing value (see Features) to testX var
func youHaveNumber(arg1 int) error {
	testX = arg1
	return nil
}

// Use the AddTwoNumbers functon to add the testing values
func youAddToThisNumber(arg1 int) error {
	testY = arg1
	testResult = AddTwoNumbers(testX, testY)
	return nil
}

// Test if the AddTwoNumbers function's result equals the testing value (see Features)
func theSummaryResultShouldBe(arg1 int) error {
	if testResult != arg1 {
		return fmt.Errorf("There's an error. %d + %d returned %d instead of %d", testX, testY, testResult, arg1)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^you have number (\d+)$`, youHaveNumber)
	s.Step(`^you add (\d+) to this number$`, youAddToThisNumber)
	s.Step(`^the summary result should be (\d+)$`, theSummaryResultShouldBe)
}
