package main

import "testing"

func TestInvalidYear(t *testing.T){

	days, err :=calculateDuration("01/12/2010", "01/12/2001")
	if err == nil {
		t.Error("Year validation is wrong")
	} 
}

func TestInvalidMonth(t *testing.T){

	days, err :=calculateDuration("01/12/2010", "01/11/2010")
	if err == nil {
		t.Error("Month validation is wrong")
	}
}

func TestInvalidDay(t *testing.T){

	days, err :=calculateDuration("31/12/2010", "01/12/2010")
	if err == nil {
		t.Error("Day validation is wrong")
	}
}

func TestInvalidFormat(t *testing.T){

	days, err :=calculateDuration("01/Dec/2010", "01/12/2010")
	if err == nil {
		t.Error("Format validation is wrong")
	}
}