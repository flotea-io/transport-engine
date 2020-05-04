package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type GtfsIndexes_20190416_143629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GtfsIndexes_20190416_143629{}
	m.Created = "20190416_143629"

	migration.Register("GtfsIndexes_20190416_143629", m)
}

// Run the migrations
func (m *GtfsIndexes_20190416_143629) Up() {

	//m.SQL("ALTER TABLE gtfs_agency ADD CONSTRAINT agency_name_pkey PRIMARY KEY (agency_id);");
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_name SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_url SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_timezone SET NOT NULL;")

	//m.SQL("ALTER TABLE gtfs_stops ADD CONSTRAINT stops_id_pkey PRIMARY KEY (stop_id);");
	//m.SQL("ALTER TABLE gtfs_stops ALTER COLUMN stop_name SET NOT NULL;");
	//m.SQL("ALTER TABLE gtfs_stops ADD CONSTRAINT stop_location_fkey FOREIGN KEY (location_type) REFERENCES gtfs_location_types(location_type);");
	//m.SQL("ALTER TABLE gtfs_stops ADD CONSTRAINT stop_parent_fkey FOREIGN KEY (parent_station) REFERENCES gtfs_stops(stop_id);");

	//m.SQL("ALTER TABLE gtfs_routes ADD CONSTRAINT routes_id_pkey PRIMARY KEY (route_id);");
	//m.SQL("ALTER TABLE gtfs_routes ADD CONSTRAINT routes_agency_fkey FOREIGN KEY (agency_id) REFERENCES gtfs_agency(agency_id);");
	//m.SQL("ALTER TABLE gtfs_routes ADD CONSTRAINT routes_rtype_fkey FOREIGN KEY (route_type) REFERENCES gtfs_route_types(route_type);");

	//m.SQL("ALTER TABLE gtfs_calendar ADD CONSTRAINT calendar_sid_pkey PRIMARY KEY (service_id);");
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN monday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN tuesday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN wednesday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN thursday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN friday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN saturday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN sunday SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN start_date SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN end_date SET NOT NULL;")

	//--ALTER TABLE gtfs_calendar_dates ADD CONSTRAINT cal_sid_fkey
	//--      FOREIGN KEY (service_id)
	//--      REFERENCES gtfs_calendar(service_id);

	//m.SQL("ALTER TABLE gtfs_fare_attributes ADD CONSTRAINT fare_id_pkey PRIMARY KEY (fare_id);");
	m.SQL("ALTER TABLE gtfs_fare_attributes ALTER COLUMN price SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_fare_attributes ALTER COLUMN currency_type SET NOT NULL;")
	//m.SQL("ALTER TABLE gtfs_fare_attributes ADD CONSTRAINT fare_pay_fkey FOREIGN KEY (payment_method) REFERENCES gtfs_payment_methods(payment_method);");
	//m.SQL("ALTER TABLE gtfs_fare_attributes ADD CONSTRAINT fare_agency_fkey FOREIGN KEY (agency_id) REFERENCES gtfs_agency(agency_id);");

	//m.SQL("ALTER TABLE gtfs_fare_rules ADD CONSTRAINT farer_id_pkey FOREIGN KEY (fare_id) REFERENCES gtfs_fare_attributes(fare_id);");
	//m.SQL("ALTER TABLE gtfs_fare_rules ADD CONSTRAINT fare_rid_fkey FOREIGN KEY (route_id) REFERENCES gtfs_routes(route_id);");

	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_id SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_lat SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_lon SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_sequence SET NOT NULL;")

	//m.SQL("ALTER TABLE gtfs_trips ADD CONSTRAINT trip_id_pkey PRIMARY KEY (trip_id);");
	//m.SQL("ALTER TABLE gtfs_trips ADD CONSTRAINT trip_rid_fkey FOREIGN KEY (route_id) REFERENCES gtfs_routes(route_id);");

	//--ALTER TABLE gtfs_trips ADD CONSTRAINT trip_sid_fkey
	//--      FOREIGN KEY (service_id)
	//--      REFERENCES gtfs_calendar(service_id);

	//m.SQL("ALTER TABLE gtfs_trips ADD CONSTRAINT trip_did_fkey FOREIGN KEY (direction_id) REFERENCES gtfs_directions(direction_id);");
	m.SQL("ALTER TABLE gtfs_trips ALTER COLUMN direction_id SET NOT NULL;")

	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_tid_fkey FOREIGN KEY (trip_id) REFERENCES gtfs_trips(trip_id);");
	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_sid_fkey FOREIGN KEY (stop_id) REFERENCES gtfs_stops(stop_id);");
	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_ptype_fkey FOREIGN KEY (pickup_type) REFERENCES gtfs_pickup_dropoff_types(type_id);");
	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_dtype_fkey FOREIGN KEY (drop_off_type) REFERENCES gtfs_pickup_dropoff_types(type_id);");
	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_arrtime_check CHECK (arrival_time LIKE '__:__:__');");
	//m.SQL("ALTER TABLE gtfs_stop_times ADD CONSTRAINT times_deptime_check CHECK (departure_time LIKE '__:__:__');");
	m.SQL("ALTER TABLE gtfs_stop_times ALTER COLUMN stop_sequence SET NOT NULL;")

	//m.SQL("create index arr_time_index on gtfs_stop_times(arrival_time_seconds);");
	//m.SQL("create index dep_time_index on gtfs_stop_times(departure_time_seconds);");
	//m.SQL("create index stop_seq_index on gtfs_stop_times(trip_id,stop_sequence);");

	//m.SQL("ALTER TABLE gtfs_frequencies ADD CONSTRAINT freq_tid_fkey FOREIGN KEY (trip_id) REFERENCES gtfs_trips(trip_id);");
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN start_time SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN end_time SET NOT NULL;")
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN headway_secs SET NOT NULL;")

	//m.SQL("ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_fsid_fkey FOREIGN KEY (from_stop_id) REFERENCES gtfs_stops(stop_id);");
	//m.SQL("ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_tsid_fkey FOREIGN KEY (to_stop_id) REFERENCES gtfs_stops(stop_id);");
	//m.SQL("ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_xt_fkey FOREIGN KEY (transfer_type) REFERENCES gtfs_transfer_types(transfer_type);");
	//m.SQL("ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_frid_fkey FOREIGN KEY (from_route_id) REFERENCES gtfs_routes(route_id);");
	//m.SQL("ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_trid_fkey FOREIGN KEY (to_route_id) REFERENCES gtfs_routes(route_id);");

	//--ALTER TABLE gtfs_transfers ADD CONSTRAINT xfer_sid_fkey
	//--      FOREIGN KEY (service_id)
	//--      REFERENCES gtfs_calendar(service_id);

}

// Reverse the migrations
func (m *GtfsIndexes_20190416_143629) Down() {

	m.SQL("drop index IF EXISTS arr_time_index;")
	m.SQL("drop index IF EXISTS dep_time_index;")
	m.SQL("drop index IF EXISTS stop_seq_index;")

	m.SQL("ALTER TABLE gtfs_agency DROP CONSTRAINT IF EXISTS agency_name_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stops DROP CONSTRAINT IF EXISTS stops_id_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stops DROP CONSTRAINT IF EXISTS stop_location_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stops DROP CONSTRAINT IF EXISTS stop_parent_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_routes DROP CONSTRAINT IF EXISTS routes_id_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_routes DROP CONSTRAINT IF EXISTS routes_agency_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_routes DROP CONSTRAINT IF EXISTS routes_rtype_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_calendar DROP CONSTRAINT IF EXISTS calendar_sid_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_calendar_dates DROP CONSTRAINT IF EXISTS cal_sid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_fare_attributes DROP CONSTRAINT IF EXISTS fare_id_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_fare_attributes DROP CONSTRAINT IF EXISTS fare_pay_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_fare_attributes DROP CONSTRAINT IF EXISTS fare_agency_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_fare_rules DROP CONSTRAINT IF EXISTS fare_rid_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_fare_rules DROP CONSTRAINT IF EXISTS fare_rid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_shapes DROP CONSTRAINT IF EXISTS shape_shape_constr;")
	m.SQL("ALTER TABLE gtfs_trips DROP CONSTRAINT IF EXISTS trip_id_pkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_trips DROP CONSTRAINT IF EXISTS trip_rid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_trips DROP CONSTRAINT IF EXISTS trip_sid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_trips DROP CONSTRAINT IF EXISTS trip_did_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_tid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_sid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_ptype_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_dtype_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_arrtime_check;")
	m.SQL("ALTER TABLE gtfs_frequencies DROP CONSTRAINT IF EXISTS freq_tid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_stop_times DROP CONSTRAINT IF EXISTS times_deptime_check;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_fsid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_tsid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_xt_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_frid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_trid_fkey CASCADE;")
	m.SQL("ALTER TABLE gtfs_transfers DROP CONSTRAINT IF EXISTS xfer_sid_fkey CASCADE;")

	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_name DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_url DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_agency ALTER COLUMN agency_timezone DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_stops ALTER COLUMN stop_name DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN monday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN tuesday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN wednesday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN thursday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN friday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN saturday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN sunday DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN start_date DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_calendar ALTER COLUMN end_date DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_fare_attributes ALTER COLUMN price DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_fare_attributes ALTER COLUMN currency_type DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_id DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_lat DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_lon DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_shapes ALTER COLUMN shape_pt_sequence DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_trips ALTER COLUMN direction_id DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_stop_times ALTER COLUMN stop_sequence DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN start_time DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN end_time DROP NOT NULL;")
	m.SQL("ALTER TABLE gtfs_frequencies ALTER COLUMN headway_secs DROP NOT NULL;")

}
