package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type RoutesTickets struct {
	TicketId    int    `orm:"column(ticket_id);pk"`
	RouteId     int    `orm:"column(route_id);"`
	Time        int    `orm:"column(time);null"`
	BuyerWallet string `orm:"column(buyer_wallet);null"`
}

func init() {
	orm.RegisterModel(new(RoutesTickets))
}

func (u *RoutesTickets) TableName() string {
	return "routes_tickets"
}

func AddRoutesTickets(m *RoutesTickets) {
	o := orm.NewOrm()
	//orm.Debug = true
	_, err := o.Insert(m)
	if err != nil {
		fmt.Println(err.Error())
	}
}
