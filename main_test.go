package main

import "testing"

//inputs to test the validateDate function
var testDates = []struct {
	dates    []string
	expected int
}{
	//Valid Dates
	{[]string{"01", "01", "1904"}, 1},
	{[]string{"29", "02", "2020"}, 1},
	{[]string{"31", "12", "2499"}, 1},

	//Invalid Dates
	{[]string{"56", "02", "2001"}, -2},
	{[]string{"01", "13", "2020"}, -2},
	{[]string{"29", "02", "2021"}, -2},
	{[]string{"15", "04", "3000"}, -2},
}

//inputs to test the countDays function
var testCountDates = []struct {
	date1    []string
	date2    []string
	expected int
}{
	{[]string{"2", "6", "1983"}, []string{"22", "6", "1983"}, 19},
	{[]string{"4", "7", "1984"}, []string{"25", "12", "1984"}, 173},
	{[]string{"3", "1", "1989"}, []string{"3", "8", "1983"}, 1979},
}

//Unit Test Case to test the validateDate function
func TestValidateDate(t *testing.T) {

	for _, tc := range testDates {
		got := validateDate(tc.dates)
		if tc.expected != got {
			t.Errorf("Expected '%d', but got '%d'", tc.expected, got)
		}
	}
}

//Unit Test Case to test the countDays function
func TestCountDays(t *testing.T) {

	for _, tc := range testCountDates {
		got := countDays(tc.date1, tc.date2)
		if tc.expected != got {
			t.Errorf("Expected '%d', but got '%d'", tc.expected, got)
		}
	}
}
