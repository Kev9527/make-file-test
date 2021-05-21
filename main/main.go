package main

import (
	"flag"
	"fmt"
	"time"
)

var timeZone = flag.String("zone", timeZoneDefault, "时区")

const (
	timeZoneDefault = "Asia/Shanghai"
)

func main() {
	flag.Parse()
	t := time.Now().UTC()
	locZone := getLocalTimeZone(*timeZone)
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
