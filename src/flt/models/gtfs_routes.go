/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"github.com/astaxie/beego/orm"
)

type ScheduleMap map[string]interface{}

type GtfsRoutes struct {
	RouteId        int    `orm:"column(route_id);null;pk"`
	AgencyId       int    `orm:"column(agency_id);null"`
	RouteShortName string `orm:"column(route_short_name);null"`
	RouteLongName  string `orm:"column(route_long_name);null"`
	RouteDesc      string `orm:"column(route_desc);null"`
	RouteType      int    `orm:"column(route_type);null"`
	RouteUrl       string `orm:"column(route_url);null"`
	RouteColor     string `orm:"column(route_color);null"`
	RouteTextColor string `orm:"column(route_text_color);null"`
	TripWallet     string `orm:"column(trip_wallet);null"`
	Places         int    `orm:"column(places);null"`
	Schedule       string `orm:"column(schedule);null"`
	Enabled        bool   `orm:"column(enabled);null"`
}

func (u *GtfsRoutes) TableName() string {
	return "gtfs_routes"
}

func init() {
	orm.RegisterModel(new(GtfsRoutes))
}

func AddRoute(m *GtfsRoutes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateRouteById(m *GtfsRoutes) (err error) {
	o := orm.NewOrm()
	v := GtfsRoutes{RouteId: m.RouteId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			_ = num
			//fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
