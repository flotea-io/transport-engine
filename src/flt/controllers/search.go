/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
	"flt/models"
	"fmt"
	"strconv"
	"strings"

	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// SearchController operations for Search
type SearchController struct {
	beego.Controller
}

type jsonStruct struct {
	Data string `json:"data"`
}

type returnJson struct {
	Length int
}

// GetOne ...
// @Title GetOne
// @Description get Search by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Search
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SearchController) Get() {

	var uid int64
	c.Ctx.Input.Bind(&uid, "id") //id ==123

	if uid != -1 {
		agency, err := models.GetAgencyById(uid)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = agency
		}
	}
	c.ServeJSON()
}

type SearchQuery struct {
	FromLat string
	FromLng string
	ToLat   string
	ToLng   string
	Radius  int // in meters
	Date    []int
	Limit   int
}

type RoutesStruct struct {
	RouteId      int     `orm:"column(route_id);null;pk"`
	RouteDesc    string  `orm:"column(route_desc);null"`
	RouteType    int     `orm:"column(route_type);null"`
	TripWallet   string  `orm:"column(trip_wallet);null"`
	Places       int     `orm:"column(places);null"`
	Schedule     string  `orm:"column(schedule);null"`
	CurrencyType string  `orm:"column(currency_type)"`
	Price        float64 `orm:"column(price)"`

	FromLabel string  `orm:"column(from_label);null"`
	ToLabel   string  `orm:"column(to_label);null"`
	FromLat   float32 `orm:"column(from_lat);null"`
	FromLng   float32 `orm:"column(from_lng);null"`
	ToLat     float32 `orm:"column(to_lat);null"`
	ToLng     float32 `orm:"column(to_lng);null"`

	AgencyId   int    `orm:"column(agency_id);null;pk"`
	AgencyName string `orm:"column(agency_name)"`
	AgencyUrl  string `orm:"column(agency_url);null"`
}

type UserTickets struct {
	Count   int
	Time    int
	RouteId int
}

// GetAll ...
// @Title GetAll
// @Description get Search
// @Param   brand_id    query   string true "Latitude from"
// @Failure 403
// @router /
func (c *SearchController) GetAll() {
	sq := SearchQuery{}
	c.ParseForm(&sq)
	c.Ctx.Input.Bind(&sq.Date, "Date")

	if sq.Radius == 0 {
		sq.Radius = 10 * 1000 // 10 Km
	}

	if sq.Limit == 0 || sq.Limit > 50 {
		sq.Limit = 10 // Default 10 records, max 50
	}

	sql := ""
	if len(sq.Date) == 3 {
		sql += " AND IsOpen(schedule::TEXT, ARRAY" + strings.Replace(fmt.Sprint(sq.Date), " ", ",", -1) + ")"
	}

	if len(sq.FromLat) > 0 && len(sq.FromLng) > 0 {
		sql += " AND ST_DWithin(from_geom, ST_MakePoint(" + sq.FromLng + "," + sq.FromLat + ")::geography, " + strconv.Itoa(sq.Radius) + ")"
	}
	if len(sq.ToLat) > 0 && len(sq.ToLng) > 0 {
		sql += " AND ST_DWithin(to_geom, ST_MakePoint(" + sq.ToLng + "," + sq.ToLat + ")::geography, " + strconv.Itoa(sq.Radius) + ")"
	}

	sql = "SELECT gtfs_routes.route_id, route_desc, route_type, trip_wallet, places, " +
		"gtfs_routes.agency_id, agency_name, agency_url, price, currency_type, " +
		"schedule, from_label, to_label, ST_Y(from_geom) as from_lat, ST_X(from_geom) as from_lng, ST_Y(to_geom) as to_lat, ST_X(to_geom) as to_lng " +
		"FROM gtfs_routes LEFT JOIN gtfs_agency ON gtfs_routes.agency_id = gtfs_agency.agency_id " +
		"LEFT JOIN gtfs_fare_rules ON gtfs_routes.route_id = gtfs_fare_rules.route_id " +
		"LEFT JOIN gtfs_fare_attributes ON gtfs_fare_rules.fare_id = gtfs_fare_attributes.fare_id " +
		"LEFT JOIN routes_stations ON gtfs_routes.route_id = routes_stations.route_id " +
		"WHERE enabled = true" + sql + " LIMIT " + strconv.Itoa(sq.Limit)

	var routes = []RoutesStruct{}

	o := orm.NewOrm()
	num, err := o.Raw(sql).QueryRows(&routes)
	if err == nil {
		_ = num
		//fmt.Println("Routes: ", num)
	} else {
		fmt.Println(err.Error())
	}
	//fmt.Println("request data", sql)

	c.Data["json"] = routes //returnJson{Length: 3}
	c.ServeJSON()
}

// @Param	route_id	path 	int	true "Route Id"
// @Param	time	path 	int	true "Time"
// @router /buyed-tickets/:route_id/:time [get]
func (c *SearchController) BuyedTickets() {
	type Tickets struct {
		Count   int
		Tickets []models.RoutesTickets
	}
	sql := "SELECT * FROM routes_tickets WHERE route_id = " + c.GetString(":route_id") +
		" AND time = " + c.GetString(":time")

	var tickets = []models.RoutesTickets{}

	o := orm.NewOrm()
	num, err := o.Raw(sql).QueryRows(&tickets)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.Data["json"] = Tickets{
		Count:   int(num),
		Tickets: tickets,
	}
	c.ServeJSON()
}

type TimesResult struct {
	Time  int
	Count int
}

// @Param	route_id	path 	int	true "Route Id"
// @Param	times	query 	string	true "Array of times"
// @router /buyed-tickets-in-times/:route_id [post]
func (c *SearchController) BuyedTicketsInTimes() {
	type Tickets struct {
		Count   int
		Tickets []models.RoutesTickets
	}
	var times = c.GetString("times")

	sql := "SELECT time, COUNT(time) FROM routes_tickets WHERE route_id = " + c.GetString(":route_id")
	if times != "" {
		sql += " AND time IN (" + times + ")"
	}
	sql += " GROUP BY time"
	//fmt.Println(sql)
	var tickets = []TimesResult{}

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&tickets)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp := make(map[int]int)

	for i := 0; i < len(tickets); i++ {
		resp[tickets[i].Time] = tickets[i].Count
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// @Param	wallet	query 	string	true "Wallet address"
// @router /passager-tickets [post]
func (c *SearchController) PassagerTickets() {

	sql := "SELECT count(time), time, route_id FROM routes_tickets WHERE buyer_wallet = '" + c.GetString("wallet") +
		"' GROUP BY(time, buyer_wallet, route_id) ORDER BY time DESC"

	var tickets = []UserTickets{}

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&tickets)
	if err != nil {
		fmt.Println(err.Error())
	}

	var Ids = []int{}

	for i := 0; i < len(tickets); i++ {
		if !intExist(Ids, tickets[i].RouteId) {
			Ids = append(Ids, tickets[i].RouteId)
		}
	}

	sql = "SELECT gtfs_routes.route_id, route_desc, route_type, trip_wallet, places, " +
		"gtfs_routes.agency_id, agency_name, agency_url, price, currency_type, " +
		"schedule, from_label, to_label, ST_Y(from_geom) as from_lat, ST_X(from_geom) as from_lng, ST_Y(to_geom) as to_lat, ST_X(to_geom) as to_lng " +
		"FROM gtfs_routes LEFT JOIN gtfs_agency ON gtfs_routes.agency_id = gtfs_agency.agency_id " +
		"LEFT JOIN gtfs_fare_rules ON gtfs_routes.route_id = gtfs_fare_rules.route_id " +
		"LEFT JOIN gtfs_fare_attributes ON gtfs_fare_rules.fare_id = gtfs_fare_attributes.fare_id " +
		"LEFT JOIN routes_stations ON gtfs_routes.route_id = routes_stations.route_id " +
		"WHERE gtfs_routes.route_id IN (" + strings.Trim(strings.Replace(fmt.Sprint(Ids), " ", ",", -1), "[]") + ")"

	var routes = []RoutesStruct{}

	_, err = o.Raw(sql).QueryRows(&routes)
	if err != nil {
		fmt.Println(err.Error())
	}

	type TicketsRoutesStructs struct {
		Tickets []UserTickets
		Routes  []RoutesStruct
	}

	c.Data["json"] = TicketsRoutesStructs{tickets, routes}
	c.ServeJSON()
}

// @Param	wallet	query 	string	true "Carrier wallet address"
// @router /carrier-routes [post]
func (c *SearchController) CarrierRoutes() {
	sql := "SELECT gtfs_routes.route_id, route_desc, route_type, trip_wallet, places, " +
		"gtfs_routes.agency_id, agency_name, agency_url, price, currency_type, " +
		"schedule, from_label, to_label, ST_Y(from_geom) as from_lat, ST_X(from_geom) as from_lng, ST_Y(to_geom) as to_lat, ST_X(to_geom) as to_lng " +
		"FROM gtfs_routes LEFT JOIN gtfs_agency ON gtfs_routes.agency_id = gtfs_agency.agency_id " +
		"LEFT JOIN gtfs_fare_rules ON gtfs_routes.route_id = gtfs_fare_rules.route_id " +
		"LEFT JOIN gtfs_fare_attributes ON gtfs_fare_rules.fare_id = gtfs_fare_attributes.fare_id " +
		"LEFT JOIN routes_stations ON gtfs_routes.route_id = routes_stations.route_id " +
		"WHERE gtfs_agency.agency_wallet ='" + c.GetString("wallet") + "'"

	var routes = []RoutesStruct{}
	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&routes)
	if err != nil {
		fmt.Println(err.Error())
	}

	var Ids = []int{}

	for i := 0; i < len(routes); i++ {
		if !intExist(Ids, routes[i].RouteId) {
			Ids = append(Ids, routes[i].RouteId)
		}
	}

	sql = "SELECT * FROM routes_tickets WHERE route_id IN (" +
		strings.Trim(strings.Replace(fmt.Sprint(Ids), " ", ",", -1), "[]") + ")" +
		" ORDER BY time ASC, ticket_id ASC"

	var tickets = []models.RoutesTickets{}

	_, err = o.Raw(sql).QueryRows(&tickets)
	if err != nil {
		fmt.Println(err.Error())
	}

	type TicktetInTime struct {
		TicketId    int
		BuyerWallet string
	}
	var t = make(map[int]map[int][]TicktetInTime)

	for i := 0; i < len(tickets); i++ {
		var tit = TicktetInTime{tickets[i].TicketId, tickets[i].BuyerWallet}

		if t[tickets[i].RouteId] == nil {
			t[tickets[i].RouteId] = make(map[int][]TicktetInTime)
		}

		var r = t[tickets[i].RouteId][tickets[i].Time]
		t[tickets[i].RouteId][tickets[i].Time] = append(r, tit)
	}

	type TicketsRoutesStructs struct {
		Tickets map[int]map[int][]TicktetInTime
		Routes  []RoutesStruct
	}
	c.Data["json"] = TicketsRoutesStructs{Tickets: t, Routes: routes}
	c.ServeJSON()
}

func intExist(slice []int, item int) bool {
	for i, _ := range slice {
		if slice[i] == item {
			return true
		}
	}
	return false
}
