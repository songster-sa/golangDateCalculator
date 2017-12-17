package main

import (
"strings"
"strconv"
)

func getDaysLeftInMonth (day int, month int, year int) (toRet int){

	toRet = 0

	days31:=[]int{1, 3, 5, 7, 8, 10, 12}
	days30:=[]int{4,6,9,11}

	if contains(days31,month) {
		toRet = 31 - day
	} else if contains(days30,month) {
		toRet = 30 - day
	} else if month==2 {

		if isLeapYear(year){
			toRet = 29 - day
		}else{
			toRet = 28 - day
		}
	}

	return toRet
}

func getWholeMonthsInDays (from int, to int, year int, inclFrom bool, inclTo bool) (toRet int) {
	toRet=0

	days31:=[]int{1, 3, 5, 7, 8, 10, 12}
	days30:=[]int{4,6,9,11}

	month31:=0
	month30:=0
	feb:=0

	i:=from
	j:=to

	if !inclFrom {
		i++
	}
	if !inclTo {
		j--
	}

	for ;i <=j;i++ {

		if contains(days31,i) {
			month31++
		} else if contains(days30,i) {
			month30++
		} else if i==2 {
			feb++
		}
	}

	toRet = month31 * 31
	toRet = toRet + (month30 * 30)

	if isLeapYear(year) {
		toRet = toRet + (feb * 29)
	} else {
		toRet = toRet + (feb * 28)
	}

	return toRet
}

func contains (input []int, val int) (bool){
	for i:=range input {
		if input[i]==val {
			return true;
		}
	}
	return false;
}

func getPartsOfDate (date string) ( arr [3]int, err error ){

	parts:=strings.Split(date,"/")

	i,err := strconv.Atoi(parts[0])
	arr[0]=i

	i,err = strconv.Atoi(parts[1])
	arr[1]=i

	i,err = strconv.Atoi(parts[2])
	arr[2]=i

	return arr, err

}

func isLeapYear (year int) (bool){
	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		return true;
	}
	return false;
}