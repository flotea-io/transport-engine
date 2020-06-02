/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ChangesForBlockchain_20190626_130621 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ChangesForBlockchain_20190626_130621{}
	m.Created = "20190626_130621"

	migration.Register("ChangesForBlockchain_20190626_130621", m)
}

// Run the migrations
func (m *ChangesForBlockchain_20190626_130621) Up() {
	m.SQL("ALTER TABLE gtfs_agency ADD COLUMN agency_wallet text;")
	m.SQL("ALTER TABLE gtfs_agency ADD COLUMN agency_address text;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_url DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_timezone DROP NOT NULL;")

	m.SQL("ALTER TABLE gtfs_routes ADD COLUMN trip_wallet text;")
	m.SQL("ALTER TABLE gtfs_routes ADD COLUMN places int;")
	m.SQL("ALTER TABLE gtfs_routes ADD COLUMN schedule jsonb;")
	m.SQL("ALTER TABLE gtfs_routes ADD COLUMN enabled bool;")

	m.SQL("create table routes_stations (route_id int UNIQUE REFERENCES gtfs_routes(route_id) ON DELETE CASCADE, from_label text, to_label text);")
	m.SQL("SELECT AddGeometryColumn('routes_stations', 'from_geom', 4326, 'POINT', 2);")
	m.SQL("SELECT AddGeometryColumn('routes_stations', 'to_geom', 4326, 'POINT', 2);")

	m.SQL("create table routes_tickets (ticket_id int, route_id int REFERENCES gtfs_routes(route_id) ON DELETE CASCADE, time int, buyer_wallet text, refunded bool);")
	m.SQL("CREATE INDEX ON routes_tickets (ticket_id);")
}

// Reverse the migrations
func (m *ChangesForBlockchain_20190626_130621) Down() {

	m.SQL("DROP TABLE IF EXISTS routes_tickets;")
	m.SQL("DROP TABLE IF EXISTS routes_stations;")

	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN to_geom;")
	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN from_geom;")
	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN enabled;")
	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN schedule;")
	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN places;")
	m.SQL("ALTER TABLE gtfs_routes DROP COLUMN trip_wallet;")

	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_timezone SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_url SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency DROP COLUMN agency_address;")
	m.SQL("ALTER TABLE gtfs_agency DROP COLUMN agency_wallet;")
}
