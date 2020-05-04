package main

import (
	timespan "flt/utils/timespan"
)

func IsOpen(json string, event []int) bool {
	return timespan.IsOpen(json, event)
}

func StartInDay(schedule string, day []int) bool {
	return timespan.StartInDay(schedule, day)
}
