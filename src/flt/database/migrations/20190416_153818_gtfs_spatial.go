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
type GtfsSpatial_20190416_153818 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GtfsSpatial_20190416_153818{}
	m.Created = "20190416_153818"

	migration.Register("GtfsSpatial_20190416_153818", m)
}

// Run the migrations
func (m *GtfsSpatial_20190416_153818) Up() {

	m.SQL("SELECT AddGeometryColumn('gtfs_stops', 'loc_geom', 4326, 'POINT', 2);");

	// Create spatial index
	m.SQL("CREATE INDEX gtfs_stops_the_geom_gist ON gtfs_stops using gist (loc_geom gist_geometry_ops_2d);")

  /* THIS MAYBE USED WHEN WE NEED TO GENERATE SHAPES FROM POINTS */
	//m.SQL("CREATE INDEX 'gtfs_stops_the_geom_gist' ON 'gtfs_stops' using gist ('the_geom' gist_geometry_ops_2d);")

	/* @TODO We need think about spatial on our project

	// Create new table to store the shape geometries
	m.SQL("CREATE TABLE gtfs_shape_geoms (shape_id text);")

	// Add the_geom column to the gtfs_shape_geoms table - a 2D linestring geometry
	m.SQL("SELECT AddGeometryColumn('gtfs_shape_geoms', 'the_geom', 4326, 'LINESTRING', 2);")

	INSERT INTO gtfs_shape_geoms
	SELECT shape.shape_id, ST_SetSRID(ST_MakeLine(shape.the_geom), 4326) As new_geom
	  FROM (
	    SELECT shape_id, ST_MakePoint(shape_pt_lon, shape_pt_lat) AS the_geom
	    FROM gtfs_shapes
	    ORDER BY shape_id, shape_pt_sequence
	  ) AS shape
	GROUP BY shape.shape_id;

	-- Create spatial index
	CREATE INDEX "gtfs_shape_geoms_the_geom_gist" ON "gtfs_shape_geoms" using gist ("the_geom" gist_geometry_ops_2d);

	COMMIT;
	*/
}

// Reverse the migrations
func (m *GtfsSpatial_20190416_153818) Down() {
	m.SQL("drop index gtfs_stops_the_geom_gist;")
	m.SQL("ALTER TABLE gtfs_stops DROP COLUMN loc_geom;");
}
