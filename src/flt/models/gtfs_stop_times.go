package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type GtfsStopTimes struct {
	TripId               int     `orm:"column(trip_id);null"`
	ArrivalTime          string  `orm:"column(arrival_time);null"`
	DepartureTime        string  `orm:"column(departure_time);null"`
	StopId               int     `orm:"column(stop_id);pk;null"`
	StopSequence         int     `orm:"column(stop_sequence)"`
	StopHeadsign         string  `orm:"column(stop_headsign);null"`
	PickupType           int     `orm:"column(pickup_type);null"`
	DropOffType          int     `orm:"column(drop_off_type);null"`
	ShapeDistTraveled    float64 `orm:"column(shape_dist_traveled);null"`
	ArrivalTimeSeconds   int     `orm:"column(arrival_time_seconds);null"`
	DepartureTimeSeconds int     `orm:"column(departure_time_seconds);null"`
}

func (u *GtfsStopTimes) TableName() string {
	return "gtfs_stop_times"
}

func init() {
	orm.RegisterModel(new(GtfsStopTimes))
}

func AddStopTime(idStop int64, tripId int64, seq int, arrivalTime string, departureTime string) {
	o := orm.NewOrm()
	var stopTime GtfsStopTimes
	stopTime.TripId = int(tripId)
	stopTime.StopId = int(idStop)
	stopTime.ArrivalTime = arrivalTime
	stopTime.DepartureTime = departureTime
	stopTime.StopSequence = seq

	_, err := o.Insert(&stopTime)
	if err != nil {
		fmt.Println(err)
	}
}
