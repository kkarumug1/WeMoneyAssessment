package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	InputDateMessage1 string `json:"inputdate1_message"`
	InputDateMessage2 string `json:"inputdate2_message"`
	MinimumYear       int    `json:"minimumYear"`
	MaximumYear       int    `json:"maximumYear"`
	InputReadError    string `json:"input_read_error"`
	InvalidDateFormat string `json:"invalid_date_format_message"`
	InvalidDate       string `json:"invalid_date_message"`
}

//function to check the validity of the day and month combination for a given year
func checkDayMonthCombination(day int, month int, year int) bool {

	monthSet1 := []int{1, 3, 5, 7, 8, 10, 12}
	monthSet2 := []int{4, 6, 9, 11}

	if ((day >= 1 && day <= 31) && Contains(monthSet1, month)) ||
		((day >= 1 && day <= 30) && Contains(monthSet2, month)) ||
		checkFebruaryMonth(day, month, year) {
		return true
	} else {
		return false
	}
}

//function to check if a given february date is valid
func checkFebruaryMonth(day int, month int, year int) bool {
	if ((day >= 1 && day <= 28) && (month == 2)) ||
		(day == 29 && month == 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0))) {
		return true
	} else {
		return false
	}
}

//function to check if a given year is valid
func checkYear(year int, minimumYear int, maximumYear int) bool {
	if year >= minimumYear && year <= maximumYear {
		return true
	} else {
		return false
	}
}

//function to check if an element is present in the input slice
func Contains(inputSlice []int, input int) bool {
	for _, i := range inputSlice {
		if i == input {
			return true
		}
	}
	return false
}

//function to get config file
func GetConfig(path string) Configuration {
	config, error := os.Open(path)
	if error != nil {
		panic(error)
	}
	decoder := json.NewDecoder(config)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration
	}
	return configuration
}

//function to get date input
func GetInputDate(message string) []string {
	configuration := GetConfig(ConfigPath)

	fmt.Printf("%v", message)
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		panic(configuration.InputReadError)
	}
	date := strings.TrimSpace(input)
	return strings.Split(date, "/")
}

//function to parse the input date to day, month and year
func ParseDate(date []string) (int, int, int) {
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])
	return day, month, year
}

//function to get the number of leap years till the given year
func GetNumberOfLeapYears(month int, year int) int {

	if month <= 2 {
		year -= 1
	}

	ans := int(year / 4)
	ans -= int(year / 100)
	ans += int(year / 400)
	return ans
}

//function to validate the input date
func ValidateDate(date []string, MinimumYear int, MaximumYear int) bool {
	configuration := GetConfig(ConfigPath)

	if len(date) != 3 {
		panic(configuration.InvalidDateFormat)
	}

	var day, month, year = ParseDate(date)

	if checkYear(year, MinimumYear, MaximumYear) && checkDayMonthCombination(day, month, year) {
		return true
	} else {
		return false
	}
}
