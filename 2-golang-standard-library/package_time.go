package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time Local:", currentTime.Local())

	utc := time.Date(2009, time.December, 11, 23, 0, 0, 0, time.UTC)
	fmt.Println("Current Time UTC:", utc)
	fmt.Println("Current Time UTC Local:", utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("valueTime:", valueTime)
	}

	fmt.Println("valueTime.Year:", valueTime.Year())
	fmt.Println("valueTime.Month:", valueTime.Month())
	fmt.Println("valueTime.Day:", valueTime.Day())
	fmt.Println("valueTime.Hour:", valueTime.Hour())

	parse, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Parse:", parse)
	}

	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2
	fmt.Println("Duration:", duration3)
}
