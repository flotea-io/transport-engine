package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EnableExtensions_20190414_094510 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EnableExtensions_20190414_094510{}
	m.Created = "20190414_094510"

	migration.Register("EnableExtensions_20190414_094510", m)
}

// Run the migrations
func (m *EnableExtensions_20190414_094510) Up() {

	m.SQL("CREATE EXTENSION postgis;")
	m.SQL("CREATE EXTENSION postgis_topology;")
	m.SQL("CREATE EXTENSION postgis_sfcgal;")
	m.SQL("CREATE EXTENSION address_standardizer;")
	m.SQL("CREATE EXTENSION postgis_tiger_geocoder CASCADE;")
	m.SQL("CREATE EXTENSION \"uuid-ossp\";")
	m.SQL("CREATE EXTENSION pgcrypto;")

}

// Reverse the migrations
func (m *EnableExtensions_20190414_094510) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP EXTENSION postgis CASCADE;")
	m.SQL("DROP EXTENSION address_standardizer CASCADE;")
	m.SQL("DROP EXTENSION \"uuid-ossp\" CASCADE;")
	m.SQL("DROP EXTENSION pgcrypto CASCADE;")
}
