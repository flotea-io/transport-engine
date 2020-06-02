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
type GtfsTables_20190415_161354 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GtfsTables_20190415_161354{}
	m.Created = "20190415_161354"

	migration.Register("GtfsTables_20190415_161354", m)
}

// Run the migrations
func (m *GtfsTables_20190415_161354) Up() {

	m.SQL("create table gtfs_agency (agency_id int UNIQUE, agency_name text, agency_url text, agency_timezone text, agency_lang text, agency_phone text, agency_fare_url text);")

	m.SQL("create table gtfs_location_types (location_type SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_location_types(location_type, description) values (0,'stop');")
	m.SQL("insert into gtfs_location_types(location_type, description) values (1,'station');")
	m.SQL("insert into gtfs_location_types(location_type, description) values (2,'station entrance');")

	m.SQL("create table gtfs_wheelchair_boardings (wheelchair_boarding SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_wheelchair_boardings(wheelchair_boarding, description) values (0, 'No accessibility information available for the stop');")
	m.SQL("insert into gtfs_wheelchair_boardings(wheelchair_boarding, description) values (1, 'At least some vehicles at this stop can be boarded by a rider in a wheelchair');")
	m.SQL("insert into gtfs_wheelchair_boardings(wheelchair_boarding, description) values (2, 'Wheelchair boarding is not possible at this stop');")

	// related to gtfs_stops(wheelchair_accessible)
	m.SQL("create table gtfs_wheelchair_accessible (wheelchair_accessible SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_wheelchair_accessible(wheelchair_accessible, description) values (0, 'No accessibility information available for this trip');")
	m.SQL("insert into gtfs_wheelchair_accessible(wheelchair_accessible, description) values (1, 'The vehicle being used on this particular trip can accommodate at least one rider in a wheelchair');")
	m.SQL("insert into gtfs_wheelchair_accessible(wheelchair_accessible, description) values (2, 'No riders in wheelchairs can be accommodated on this trip');")

	/*
		-- unofficial features
		location_type int, --FOREIGN KEY REFERENCES gtfs_location_types(location_type)
		parent_station text, --FOREIGN KEY REFERENCES gtfs_stops(stop_id)
		stop_timezone text,
		wheelchair_boarding int --FOREIGN KEY REFERENCES gtfs_wheelchair_boardings(wheelchair_boarding)
		-- Unofficial fields

		-- select AddGeometryColumn( 'gtfs_stops', 'location', #{WGS84_LATLONG_EPSG}, 'POINT', 2 );
		-- CREATE INDEX gtfs_stops_location_ix ON gtfs_stops USING GIST ( location GIST_GEOMETRY_OPS );
	*/

	m.SQL("create table gtfs_stops (stop_id SERIAL UNIQUE, stop_name text, stop_desc text, stop_lat double precision, stop_lon double precision, zone_id text, stop_url text, stop_code text, stop_street text, stop_city text, stop_region text, stop_postcode text, stop_country text, location_type int, parent_station text, stop_timezone text, wheelchair_boarding int, direction text, position int);")
	m.SQL("create table gtfs_route_types (route_type SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_route_types (route_type, description) values (0, 'Street Level Rail');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (1, 'Underground Rail');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (2, 'Intercity Rail');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (3, 'Bus');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (4, 'Ferry');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (5, 'Cable Car');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (6, 'Suspended Car');")
	m.SQL("insert into gtfs_route_types (route_type, description) values (7, 'Steep Incline Mode');")

	m.SQL("create table gtfs_routes (route_id int UNIQUE, agency_id int REFERENCES gtfs_agency(agency_id) ON DELETE CASCADE, route_short_name text DEFAULT '', route_long_name text DEFAULT '', route_desc text, route_type int, route_url text, route_color text, route_text_color text);")

	m.SQL("create table gtfs_directions (direction_id SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_directions (direction_id, description) values (0,'This way');")
	m.SQL("insert into gtfs_directions (direction_id, description) values (1,'That way');")

	m.SQL("create table gtfs_pickup_dropoff_types (type_id SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_pickup_dropoff_types (type_id, description) values (0,'Regularly Scheduled');")
	m.SQL("insert into gtfs_pickup_dropoff_types (type_id, description) values (1,'Not available');")
	m.SQL("insert into gtfs_pickup_dropoff_types (type_id, description) values (2,'Phone arrangement only');")
	m.SQL("insert into gtfs_pickup_dropoff_types (type_id, description) values (3,'Driver arrangement only');")

	// route_id --REFERENCES gtfs_routes(route_id),
	// trip_id --REFERENCES gtfs_calendar(service_id),
	// direction_id --REFERENCES gtfs_directions(direction_id),
	// wheelchair_accessible  --FOREIGN KEY REFERENCES gtfs_wheelchair_accessible(wheelchair_accessible)
	// Unnoficial trip_type text
	m.SQL("create table gtfs_trips (route_id int REFERENCES gtfs_routes(route_id) ON DELETE CASCADE, service_id SERIAL UNIQUE, trip_id SERIAL UNIQUE, trip_headsign text, direction_id int, block_id text, shape_id text, trip_short_name text, wheelchair_accessible int, trip_type text);")

	// -- CREATE INDEX gst_trip_id_stop_sequence ON gtfs_stop_times (trip_id, stop_sequence);
	m.SQL("create table gtfs_calendar (service_id int REFERENCES gtfs_trips(service_id) ON DELETE CASCADE, monday int, tuesday int, wednesday int, thursday int, friday int, saturday int, sunday int, start_date date, end_date date);")

	// service_id REFERENCES gtfs_calendar(service_id),
	m.SQL("create table gtfs_calendar_dates (service_id int REFERENCES gtfs_trips(service_id) ON DELETE CASCADE, date date, exception_type int);")

	/* -- The following two tables are not in the spec, but they make dealing with dates and services easier
	create table service_combo_ids
	(
	combination_id serial --primary key
	);
	create table service_combinations
	(
	combination_id int , --references service_combo_ids(combination_id),
	service_id text --references gtfs_calendar(service_id)
	);*/

	m.SQL("create table gtfs_payment_methods (payment_method SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_payment_methods (payment_method, description) values (0,'On Board');")
	m.SQL("insert into gtfs_payment_methods (payment_method, description) values (1,'Prepay');")
	m.SQL("insert into gtfs_payment_methods (payment_method, description) values (2,'Advance');")

	// payment_method REFERENCES gtfs_payment_methods,
	// agency_id --REFERENCES gtfs_agency(agency_id)

	m.SQL("create table gtfs_fare_attributes (fare_id SERIAL PRIMARY KEY, price double precision, currency_type text, payment_method int, transfers int, transfer_duration int, agency_id int REFERENCES gtfs_agency(agency_id) ON DELETE CASCADE);")

	// fare_id -REFERENCES gtfs_fare_attributes(fare_id),
	// route_id -REFERENCES gtfs_routes(route_id),
	// service_id - REFERENCES gtfs_calendar(service_id) ?
	m.SQL("create table gtfs_fare_rules (fare_id int REFERENCES gtfs_fare_attributes(fare_id) ON DELETE CASCADE, route_id int REFERENCES gtfs_routes(route_id) ON DELETE CASCADE, origin_id text, destination_id text, contains_id text, service_id text);")
	m.SQL("create table gtfs_shapes (shape_id text, shape_pt_lat double precision, shape_pt_lon double precision, shape_pt_sequence int, shape_dist_traveled double precision);")

	// trip_id --REFERENCES gtfs_trips(trip_id),
	// arrival_time -- CHECK (arrival_time LIKE '__:__:__'),
	// departure_time -- CHECK (departure_time LIKE '__:__:__'),
	// stop_id  --REFERENCES gtfs_stops(stop_id),
	// pickup_type  --REFERENCES gtfs_pickup_dropoff_types(type_id),
	// drop_off_type --REFERENCES gtfs_pickup_dropoff_types(type_id),

	m.SQL("create table gtfs_stop_times (trip_id int REFERENCES gtfs_trips(trip_id) ON DELETE CASCADE, arrival_time text, departure_time text, stop_id int REFERENCES gtfs_stops(stop_id) ON DELETE CASCADE, stop_sequence int, stop_headsign text, pickup_type int, drop_off_type int, shape_dist_traveled double precision, arrival_time_seconds int, departure_time_seconds int);")

	// --create index arr_time_index on gtfs_stop_times(arrival_time_seconds);
	// --create index dep_time_index on gtfs_stop_times(departure_time_seconds);
	// -- select AddGeometryColumn( 'gtfs_shapes', 'shape', #{WGS84_LATLONG_EPSG}, 'LINESTRING', 2 );
	// trip_id  --REFERENCES gtfs_trips(trip_id),

	m.SQL("create table gtfs_frequencies (trip_id int REFERENCES gtfs_trips(trip_id) ON DELETE CASCADE, start_time text, end_time text, headway_secs int, exact_times int, start_time_seconds int, end_time_seconds int);")

	m.SQL("create table gtfs_transfer_types (transfer_type SERIAL PRIMARY KEY, description text);")

	m.SQL("insert into gtfs_transfer_types (transfer_type, description) values (0,'Preferred transfer point');")
	m.SQL("insert into gtfs_transfer_types (transfer_type, description) values (1,'Designated transfer point');")
	m.SQL("insert into gtfs_transfer_types (transfer_type, description) values (2,'Transfer possible with min_transfer_time window');")
	m.SQL("insert into gtfs_transfer_types (transfer_type, description) values (3,'Transfers forbidden');")

	// from_stop_id  --REFERENCES gtfs_stops(stop_id)
	// to_stop_id --REFERENCES gtfs_stops(stop_id)
	// transfer_type  --REFERENCES gtfs_transfer_types(transfer_type)
	// Unofficial from_route_id --REFERENCES gtfs_routes(route_id)
	// Unofficial to_route_id  --REFERENCES gtfs_routes(route_id)
	// Unofficial service_id  --REFERENCES gtfs_calendar(service_id) ?

	m.SQL("create table gtfs_transfers (from_stop_id text, to_stop_id text, transfer_type int, min_transfer_time int, from_route_id int, to_route_id text, service_id text);")
	m.SQL("create table gtfs_feed_info (feed_publisher_name text, feed_publisher_url text, feed_timezone text, feed_lang text, feed_version text, feed_start_date text, feed_end_date text);")

}

// Reverse the migrations
func (m *GtfsTables_20190415_161354) Down() {

	m.SQL("DROP TABLE IF EXISTS gtfs_agency cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_stops cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_routes cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_calendar cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_calendar_dates cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_fare_attributes cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_fare_rules cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_shapes cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_trips cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_stop_times cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_frequencies cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_transfers cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_feed_info cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_route_types cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_directions cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_pickup_dropoff_types cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_payment_methods cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_location_types cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_wheelchair_boardings cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_wheelchair_accessible cascade")
	m.SQL("DROP TABLE IF EXISTS gtfs_transfer_types cascade")

	m.SQL("DROP TABLE IF EXISTS service_combo_ids cascade")
	m.SQL("DROP TABLE IF EXISTS service_combinations cascade")

}
