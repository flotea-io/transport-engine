package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type GtfsTrips struct {
	RouteId              int    `orm:"column(route_id);null"`
	ServiceId            int    `orm:"column(service_id);null"`
	TripId               int    `orm:"column(trip_id);pk;null"`
	TripHeadsign         string `orm:"column(trip_headsign);null"`
	DirectionId          int    `orm:"column(direction_id)"`
	BlockId              string `orm:"column(block_id);null"`
	ShapeId              string `orm:"column(shape_id);null"`
	TripShortName        string `orm:"column(trip_short_name);null"`
	WheelchairAccessible int    `orm:"column(wheelchair_accessible);null"`
	TripType             string `orm:"column(trip_type);null"`
}

func (u *GtfsTrips) TableName() string {
	return "gtfs_trips"
}

func init() {
	orm.RegisterModel(new(GtfsTrips))
}

func AddTrip(routeId int) int64 {
	//orm.Debug = true
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO gtfs_trips (route_id, direction_id) VALUES ( " +
		strconv.Itoa(routeId) + ", 0 )").Exec()

	var id int64
	err = o.Raw("SELECT last_value FROM gtfs_trips_service_id_seq").QueryRow(&id)
	if err == nil {
		return id
	} else {
		return -1
	}
}
