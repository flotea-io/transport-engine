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

type GtfsFareAttributes struct {
	Id               int     `orm:"column(fare_id);pk;auto"`
	Price            float64 `orm:"column(price)"`
	CurrencyType     string  `orm:"column(currency_type)"`
	PaymentMethod    int     `orm:"column(payment_method);null"`
	Transfers        int     `orm:"column(transfers);null"`
	TransferDuration int     `orm:"column(transfer_duration);null"`
	AgencyId         int     `orm:"column(agency_id);"`
}

func (t *GtfsFareAttributes) TableName() string {
	return "gtfs_fare_attributes"
}

func init() {
	orm.RegisterModel(new(GtfsFareAttributes))
}

// AddGtfsFareAttributes insert a new GtfsFareAttributes into database and returns
// last inserted Id on success.
func AddGtfsFareAttributes(price float64, agencyId int) int {
	o := orm.NewOrm()
	var gtfsFareAttributes GtfsFareAttributes
	gtfsFareAttributes.Price = price
	gtfsFareAttributes.CurrencyType = "FLT"
	gtfsFareAttributes.PaymentMethod = 1
	gtfsFareAttributes.Transfers = 0
	gtfsFareAttributes.AgencyId = agencyId

	_, err := o.Insert(&gtfsFareAttributes)
	if err != nil {
		fmt.Println(err)
	}
	return gtfsFareAttributes.Id
}
