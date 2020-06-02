/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/tidwall/gjson"
)

type TripLocFloat struct {
	FromLat string
	FromLng string
	ToLat   string
	ToLng   string
}

type RoutesStations struct {
	RouteId   int    `orm:"column(route_id);pk"`
	FromLabel string `orm:"column(from_label);null"`
	ToLabel   string `orm:"column(to_label);null"`
	FromGeom  string `orm:"column(from_geom);null"`
	ToGeom    string `orm:"column(to_geom);null"`
}

func init() {
	orm.RegisterModel(new(RoutesStations))
}

func (u *RoutesStations) TableName() string {
	return "routes_stations"
}

func AddRouteStations(route_id int, tripLoc TripLocFloat) {
	from_label := getLocationLabel(tripLoc.FromLat, tripLoc.FromLng)
	to_label := getLocationLabel(tripLoc.ToLat, tripLoc.ToLng)
	fmt.Println(from_label, to_label, "tu")
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO routes_stations (route_id, from_label, to_label, from_geom, to_geom) VALUES ( " +
		strconv.Itoa(route_id) + ", '" + from_label + "', '" + to_label + "', " +
		"ST_SetSRID(ST_Point(" + tripLoc.FromLng + "," + tripLoc.FromLat + "),4326), " +
		"ST_SetSRID(ST_Point(" + tripLoc.ToLng + "," + tripLoc.ToLat + "),4326) )").Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getLocationLabel(lat string, lng string) string {
	resp, err := http.Get("https://geo.flotea.pl/v1/reverse?point.lat=" + lat + "&point.lon=" + lng + "&lang=EN-en")
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		value := gjson.Get(string(body), "features.0.properties.label")
		return value.String()
	}

	return ""
}
