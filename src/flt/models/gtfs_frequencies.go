/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

type GtfsFrequencies struct {
	TripId           string `orm:"column(trip_id);null"`
	StartTime        string `orm:"column(start_time)"`
	EndTime          string `orm:"column(end_time)"`
	HeadwaySecs      int    `orm:"column(headway_secs)"`
	ExactTimes       int    `orm:"column(exact_times);null"`
	StartTimeSeconds int    `orm:"column(start_time_seconds);null"`
	EndTimeSeconds   int    `orm:"column(end_time_seconds);null"`
}
