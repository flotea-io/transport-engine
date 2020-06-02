/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type GtfsFareRules struct {
	FareId        int    `orm:"column(fare_id);pk"`
	RouteId       int    `orm:"column(route_id);"`
	OriginId      string `orm:"column(origin_id);null"`
	DestinationId string `orm:"column(destination_id);null"`
	ContainsId    string `orm:"column(contains_id);null"`
	ServiceId     string `orm:"column(service_id);null"`
}

func (u *GtfsFareRules) TableName() string {
	return "gtfs_fare_rules"
}

func init() {
	orm.RegisterModel(new(GtfsFareRules))
}

// AddGtfsFareRules insert a new GtfsFareRules into database and returns
// last inserted Id on success.
func AddGtfsFareRules(fareId int, routeId int) {
	o := orm.NewOrm()
	var gtfsFareRules GtfsFareRules
	gtfsFareRules.FareId = fareId
	gtfsFareRules.RouteId = routeId

	_, err := o.Insert(&gtfsFareRules)
	if err != nil {
		fmt.Println(err)
	}
}
