/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

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
