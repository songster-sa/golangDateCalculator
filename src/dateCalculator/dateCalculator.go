package main

import (
"fmt" 
"errors"
)

func main(){

	var fromdate string
	fmt.Printf("Enter experiment beign date:\n")
	fmt.Scanln(&fromdate)

	var todate string
	fmt.Printf("Enter experiment end date:\n")
	fmt.Scanln(&todate)

	days,err := calculateDuration(fromdate, todate) 
	fmt.Println("result : ", days)
	fmt.Println("error : ", err)
}

func calculateDuration(fromdate, todate string) (toReturn int, err error){

	toReturn = 0
	err = nil
	// validate(fromdate, todate)

	from, err := getPartsOfDate(fromdate)
	to, err := getPartsOfDate(todate)

	// ----------------------------- years --------------------

	years:=0
	leapyears:=0

	 if from[2] > to[2] {
	 	err = errors.New("Invalid input : start year is after end year")
	 } else if from[2] < to[2] {
	 	years = to[2] - from[2] - 1

	 	if years >0 {
	 		for i:=from[2]+1; i<to[2];i++ {
	 			if isLeapYear(i) {
	 				leapyears++;
	 			}
	 		}
	 	}
	 }

	 toReturn = years * 365
	 toReturn = toReturn + leapyears

	 // ------------------------ months ------------

	 if from[2] == to[2] {
	 	// same year
	 	if from[1] > to[1] {
	 		err = errors.New("Invalid input : start month is after end month")
	 	} else if to[1] - from[1] >1 {
	 		daysForMonths := getWholeMonthsInDays(from[1], to[1], from[2], false, false)
	 		toReturn = toReturn + daysForMonths
	 	}
	 } else{
	 	// diferent years

	 	if from[1] != 12 {
	 		daysForMonths := getWholeMonthsInDays(from[1], 12, from[2], false, true)
	 		toReturn = toReturn + daysForMonths
	 	}

	 	if to[1] != 1 {
	 		daysForMonths := getWholeMonthsInDays(1, to[1], to[2], true, false)
	 		toReturn = toReturn + daysForMonths
	 	}
	 }

	 // ------------------ days ---------------------

	 if  from[1]==to[1] && from[2]==to[2] {
	 	// same month and year
	 	if from[0] > to[0] {
	 		err = errors.New("Invalid input : start day is after end day")
	 	} else if to[0] - from[0] > 1 {
	 		toReturn = toReturn + (to[0] - from[0] - 1)
	 	}

	 } else {
	 	// different months
	 	days := getDaysLeftInMonth(from[0],from[1],from[2])
	 	toReturn = toReturn + days

	 	toReturn = toReturn + (to[0] - 1)
	 }
	return toReturn, err

}