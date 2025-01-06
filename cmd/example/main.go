package main

import (
	"fmt"
	"time"
)

func main() {
	//loc, _ := time.LoadLocation("America/Los_Angeles")

	now := time.Now()

	// layout := "2006-01-02 15:04:05 -0800"
	layout := time.RFC3339
	nowInRFC1123 := now.Format(layout)
	dumpTime(now, "original", layout)
	//	loc, _ := time.LoadLocation("America/Los_Angeles")

	parsedTime, _ := time.Parse(layout, nowInRFC1123)
	dumpTime(parsedTime, "parsedTime", layout)
}

func dumpTime(t time.Time, msg string, layout string) {
	fmt.Printf("%s ==> time: %s; unix: %d; location: %s\n", msg, t.Format(layout), t.Unix(), t.Location().String())
}
