package main

import (
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func main() {
	fmt.Println("test")
	t := time.Now().UTC()
	locZone := getLocalTimeZone()
	parsedTime := t.In(locZone)
	fmt.Println("local time zone is: ", locZone)
	fmt.Println("parsed time is: ", parsedTime)
}

func getLocalTimeZone() *time.Location {
	local1, err1 := time.LoadLocation("") //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	return local1
}
