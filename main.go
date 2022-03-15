package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//function to parse the input date to day, month and year
func parseDate(date []string) (int, int, int) {
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])
	return day, month, year
}

//function to validate the input date
func validateDate(date []string) int {

	if len(date) != 3 {
		return -1
	}

	var day, month, year = parseDate(date)

	if year >= 1900 && year <= 2999 {
		if month >= 1 && month <= 12 {
			if (day >= 1 && day <= 31) && (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) {
				return 1
			} else if (day >= 1 && day <= 30) && (month == 4 || month == 6 || month == 9 || month == 11) {
				return 1
			} else if (day >= 1 && day <= 28) && (month == 2) {
				return 1
			} else if day == 29 && month == 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
				return 1
			} else {
				return -2
			}
		} else {
			return -2
		}
	} else {
		return -2
	}
}

//function to count the number of leap years
func numberOfLeapYears(month int, year int) int {

	if month <= 2 {
		year -= 1
	}

	ans := int(year / 4)
	ans -= int(year / 100)
	ans += int(year / 400)
	return ans
}

//function to count the number of days between two dates
func countDays(date1 []string, date2 []string) (daysCount int) {
	daysInMonths := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var day1, month1, year1 = parseDate(date1)
	var day2, month2, year2 = parseDate(date2)

	if day1+month1*100+year1*10000 > day2+month2*100+year2*10000 {
		day1, day2 = day2, day1
		month1, month2 = month2, month1
		year1, year2 = year2, year1
	}

	var n1 int = year1*365 + day1

	for i := 0; i < month1-1; i++ {
		n1 += daysInMonths[i]
	}

	n1 += numberOfLeapYears(month1, year1)

	var n2 int = year2*365 + day2

	for i := 0; i < month2-1; i++ {
		n2 += daysInMonths[i]
	}

	n2 += numberOfLeapYears(month2, year2)

	return n2 - n1 - 1
}

//driver function
func main() {

	fmt.Printf("Enter a First Date: ")
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	date1 := strings.Split(input1, "/")
	validity := validateDate(date1)
	if validity == -1 {
		fmt.Println("Error: Enter date in the following format: day/month/yearyear")
		return
	} else if validity == -2 {
		fmt.Println("Error: Invalid Date. Provid a valid Date")
		return
	}

	fmt.Println("Enter a Second Date")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	date2 := strings.Split(input2, "/")
	validity = validateDate(date2)
	if validity == -1 {
		fmt.Println("Error: Enter date in the following format: day/month/yearyear")
		return
	} else if validity == -2 {
		fmt.Println("Error: Invalid Date. Provid a valid Date")
		return
	}

	fmt.Println("The number of days between", input1, "and", input2, "is", countDays(date1, date2))

}
