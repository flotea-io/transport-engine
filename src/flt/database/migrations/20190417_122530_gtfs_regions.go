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
type GtfsRegions_20190417_122530 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GtfsRegions_20190417_122530{}
	m.Created = "20190417_122530"

	migration.Register("GtfsRegions_20190417_122530", m)
}

// Run the migrations
func (m *GtfsRegions_20190417_122530) Up() {

	// can connect a group of stations like stop_id or location.id
	m.SQL("create table gtfs_location_group_name (id SERIAL PRIMARY KEY, location_id int, name varchar(255));")
	m.SQL("create table gtfs_location_groups (id SERIAL PRIMARY KEY, locations json);")

	// Create new table to store the shape geometries
	m.SQL("CREATE TABLE gtfs_supported_areas (id SERIAL PRIMARY KEY, uuid uuid NOT NULL DEFAULT uuid_generate_v1mc(), the_geom geometry);")
	m.SQL("CREATE INDEX gtfs_supported_areas_gix_search ON gtfs_supported_areas USING GIST (the_geom);")

	//m.SQL("ALTER TABLE gtfs_calendar DROP COLUMN monday, DROP COLUMN tuesday, DROP COLUMN wednesday, DROP COLUMN thursday, DROP COLUMN friday, DROP COLUMN saturday, DROP COLUMN sunday, DROP COLUMN start_date, DROP COLUMN end_date;")
	//m.SQL("ALTER TABLE gtfs_calendar ADD COLUMN schedule json NOT NULL;")

	//m.SQL("DROP TABLE IF EXISTS gtfs_calendar_dates;")

	// Add Foreign Key for location id
	m.SQL("ALTER TABLE gtfs_location_group_name ADD CONSTRAINT group_id_pkey FOREIGN KEY (location_id) REFERENCES gtfs_location_groups(id);")

	/*
	 The detour area will be represented by a location group (a GeoJSON Polygon),
	 and the stop time will have stop_sequence=0 and to_stop_sequence with the maximal value of stop_sequence in the trip.
	*/
	//m.SQL("ALTER TABLE gtfs_stop_times ADD COLUMN to_stop_sequence int, ADD COLUMN min_arrival_time date, ADD COLUMN max_departure_time date")

	/*
		1. PK with UUID id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		2.
	*/

}

// Reverse the migrations
func (m *GtfsRegions_20190417_122530) Down() {

	m.SQL("ALTER TABLE gtfs_location_group_name DROP CONSTRAINT IF EXISTS group_id_pkey CASCADE;")

	//m.SQL("CREATE TABLE IF NOT EXISTS gtfs_calendar_dates (service_id text, date date, exception_type int);")
	//m.SQL("ALTER TABLE gtfs_calendar ADD COLUMN monday int, ADD COLUMN tuesday int, ADD COLUMN wednesday int, ADD COLUMN thursday int, ADD COLUMN friday int, ADD COLUMN saturday int, ADD COLUMN sunday int, ADD COLUMN start_date date, ADD COLUMN end_date date;")

	//m.SQL("ALTER TABLE gtfs_stop_times DROP COLUMN to_stop_sequence, DROP COLUMN min_arrival_time, DROP COLUMN max_departure_time;")

	m.SQL("DROP TABLE IF EXISTS gtfs_location_groups;")
	m.SQL("DROP TABLE IF EXISTS gtfs_location_group_name;")
	m.SQL("DROP TABLE IF EXISTS gtfs_supported_areas;")
}
