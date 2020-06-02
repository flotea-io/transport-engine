/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	//"log"
	"github.com/astaxie/beego/orm"
)

type GtfsAgency struct {
	AgencyId       int    `orm:"column(agency_id);null;pk"`
	AgencyName     string `orm:"column(agency_name)"`
	AgencyUrl      string `orm:"column(agency_url);null"`
	AgencyTimezone string `orm:"column(agency_timezone);null"`
	AgencyLang     string `orm:"column(agency_lang);null"`
	AgencyPhone    string `orm:"column(agency_phone);null"`
	AgencyFareUrl  string `orm:"column(agency_fare_url);null"`
	AgencyWallet   string `orm:"column(agency_wallet);null"`
	AgencyAddress  string `orm:"column(agency_address);null"`
}

func (u *GtfsAgency) TableName() string {
	return "gtfs_agency"
}

func init() {
	orm.RegisterModel(new(GtfsAgency))
}

// AddAgency insert a new GtfsAgency into database and returns
// last inserted Id on success.
func AddAgency(m *GtfsAgency) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAgencyById retrieves GtfsAgency by Id. Returns error if
// Id doesn't exist
func GetAgencyById(id int64) (v *GtfsAgency, err error) {
	o := orm.NewOrm()
	v = &GtfsAgency{AgencyId: int(id)}
	if err = o.QueryTable(new(GtfsAgency)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAgency retrieves all GtfsAgency matches certain condition. Returns empty list if
// no records exist
func GetAllAgency(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GtfsAgency))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []GtfsAgency
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAgency updates GtfsAgency by Id and returns error if
// the record to be updated doesn't exist
func UpdateAgencyById(m *GtfsAgency) (err error) {
	o := orm.NewOrm()
	v := GtfsAgency{AgencyId: m.AgencyId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func GetAgencyByWallet(wallet string) (m GtfsAgency, err error) {
	o := orm.NewOrm()
	v := GtfsAgency{AgencyWallet: wallet}
	err = o.Read(&v, "AgencyWallet")
	return v, err
}

func UpdateAgencyByWallet(m *GtfsAgency) (err error) {
	o := orm.NewOrm()
	agency, err := GetAgencyByWallet(m.AgencyWallet)
	if err == nil {
		var num int64
		if num, err = o.Update(agency); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// DeleteAgency deletes GtfsAgency by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAgency(id int64) (err error) {
	o := orm.NewOrm()
	v := GtfsAgency{AgencyId: int(id)}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GtfsAgency{AgencyId: int(id)}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
