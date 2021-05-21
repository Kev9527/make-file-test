package main

import (
	"flag"
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

var zone = "UTC"

func init() {
	flag.StringVar(&zone, "zone", "", "missing zoneinfo")
}

func main() {
	flag.Parse()
	fmt.Println(zone)
	t := time.Now().UTC()
	locZone := getLocalTimeZone(zone)
	parsedTime := t.In(locZone)
	fmt.Println("local time zone is: ", locZone)
	fmt.Println("parsed time is: ", parsedTime)
}

func getLocalTimeZone(locStr string) *time.Location {
	local1, err1 := time.LoadLocation(locStr) //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	return local1
}
