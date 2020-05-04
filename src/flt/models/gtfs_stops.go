package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

type GtfsStops struct {
	StopId             int     `orm:"column(stop_id);null;pk"`
	StopName           string  `orm:"column(stop_name);null"`
	StopDesc           string  `orm:"column(stop_desc);null"`
	StopLat            float64 `orm:"column(stop_lat);null"`
	StopLon            float64 `orm:"column(stop_lon);null"`
	ZoneId             string  `orm:"column(zone_id);null"`
	StopUrl            string  `orm:"column(stop_url);null"`
	StopCode           string  `orm:"column(stop_code);null"`
	StopStreet         string  `orm:"column(stop_street);null"`
	StopCity           string  `orm:"column(stop_city);null"`
	StopRegion         string  `orm:"column(stop_region);null"`
	StopPostcode       string  `orm:"column(stop_postcode);null"`
	StopCountry        string  `orm:"column(stop_country);null"`
	LocationType       int     `orm:"column(location_type);null"`
	ParentStation      string  `orm:"column(parent_station);null"`
	StopTimezone       string  `orm:"column(stop_timezone);null"`
	WheelchairBoarding int     `orm:"column(wheelchair_boarding);null"`
	Direction          string  `orm:"column(direction);null"`
	Position           int     `orm:"column(position);null"`
	LocGeom            string  `orm:"column(loc_geom);null"`
}

func (u *GtfsStops) TableName() string {
	return "gtfs_stops"
}

func init() {
	orm.RegisterModel(new(GtfsStops))
}

func AddGtfsStop(lat string, lng string) int64 {
	//orm.Debug = true
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO gtfs_stops (stop_lat, stop_lon, loc_geom) VALUES ( " +
		lat + ", " + lng + ", ST_SetSRID(ST_Point(" + lng + "," + lat + "),4326) )").Exec()

	var id int64
	err = o.Raw("SELECT last_value FROM gtfs_stops_stop_id_seq").QueryRow(&id)
	if err == nil {
		return id
	} else {
		return -1
	}
}

func RemoveStopsByIds(intIds []int64) {
	var ids = strings.Trim(strings.Replace(fmt.Sprint(intIds), " ", ",", -1), "[]")
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("DELETE FROM gtfs_stops WHERE stop_id IN (" + ids + ")")
	r.Exec()
}
