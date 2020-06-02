/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GtfsCalendar struct {
	ServiceId int       `orm:"column(service_id);pk;null"`
	Monday    int       `orm:"column(monday)"`
	Tuesday   int       `orm:"column(tuesday)"`
	Wednesday int       `orm:"column(wednesday)"`
	Thursday  int       `orm:"column(thursday)"`
	Friday    int       `orm:"column(friday)"`
	Saturday  int       `orm:"column(saturday)"`
	Sunday    int       `orm:"column(sunday)"`
	StartDate time.Time `orm:"column(start_date);type(date)"`
	EndDate   time.Time `orm:"column(end_date);type(date)"`
}

func (u *GtfsCalendar) TableName() string {
	return "gtfs_calendar"
}

func init() {
	orm.RegisterModel(new(GtfsCalendar))
}
