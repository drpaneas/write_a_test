Feature: Add 2 numbers
  In order to perform addition of 2 numbers
  As a user
  I need to be able to AddTwoNumbers the sum of them

  Scenario: Add 5 to 4
    Given you have number 5
    When you add 4 to this number
    Then the summary result should be 9

  Scenario: Add 3 to 11
    Given you have number 3
    When you add 12 to this number
    Then the summary result should be 15
