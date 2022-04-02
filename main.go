package main

import (
	"days_counter/utils"
	"fmt"
	"reflect"
	"strings"
)

//function to calculate the number of days between the two given dates
func countDays(date1 []string, date2 []string) (daysCount int) {

	if reflect.DeepEqual(date1, date2) {
		return 0
	}

	daysInMonths := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var day1, month1, year1 = utils.ParseDate(date1)
	var day2, month2, year2 = utils.ParseDate(date2)

	if day1+month1*100+year1*10000 > day2+month2*100+year2*10000 {
		day1, day2 = day2, day1
		month1, month2 = month2, month1
		year1, year2 = year2, year1
	}

	var n1 int = year1*365 + day1

	for i := 0; i < month1-1; i++ {
		n1 += daysInMonths[i]
	}

	n1 += utils.GetNumberOfLeapYears(month1, year1)

	var n2 int = year2*365 + day2

	for i := 0; i < month2-1; i++ {
		n2 += daysInMonths[i]
	}

	n2 += utils.GetNumberOfLeapYears(month2, year2)

	return n2 - n1 - 1
}

//driver function
func main() {

	//read config file
	configuration := utils.GetConfig(utils.ConfigPath)

	//process date1
	date1 := utils.GetInputDate(configuration.InputDateMessage1)
	date1Validity := utils.ValidateDate(date1, configuration.MinimumYear, configuration.MaximumYear)
	if !date1Validity {
		panic(configuration.InvalidDate)
	}

	//process date2
	date2 := utils.GetInputDate(configuration.InputDateMessage2)
	date2Validity := utils.ValidateDate(date2, configuration.MinimumYear, configuration.MaximumYear)
	if !date2Validity {
		panic(configuration.InvalidDate)
	}

	// //Print the Output
	fmt.Println("The number of days between", strings.Join(date1, "/"), "and", strings.Join(date2, "/"), "is", countDays(date1, date2))

}
