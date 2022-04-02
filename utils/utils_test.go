package utils

import (
	"testing"
)

var MinimumYear = 1900
var MaximumYear = 2999

//inputs to test the Contains function
var testContainsInput = []struct {
	inputSlice       []int
	elementToBeFound int
	expected         bool
}{
	//Valid Dates
	{[]int{2, 1, 5, 6}, 6, true},
	{[]int{2, 1, 5}, 3, false},
}

//Unit Test Case to test the Contains function
func TestContains(t *testing.T) {

	for _, tc := range testContainsInput {
		got := Contains(tc.inputSlice, tc.elementToBeFound)
		if tc.expected != got {
			t.Errorf("Expected '%t', but got '%t'", tc.expected, got)
		}
	}
}

//inputs to test the validateDate function
var testDatesInput = []struct {
	dates    []string
	expected bool
}{
	//Valid Dates
	{[]string{"01", "01", "1904"}, true},
	{[]string{"29", "02", "2020"}, true},
	{[]string{"31", "12", "2499"}, true},

	//Invalid Dates
	{[]string{"56", "02", "2001"}, false},
	{[]string{"01", "13", "2020"}, false},
	{[]string{"29", "02", "2021"}, false},
	{[]string{"15", "04", "3001"}, false},
}

//Unit Test Case to test the validateDate function
func TestValidateDate(t *testing.T) {

	for _, tc := range testDatesInput {
		got := ValidateDate(tc.dates, MinimumYear, MaximumYear)
		if tc.expected != got {
			t.Errorf("Expected '%t', but got '%t'", tc.expected, got)
		}
	}
}
