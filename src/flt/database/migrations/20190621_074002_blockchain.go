package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Blockchain_20190621_074002 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Blockchain_20190621_074002{}
	m.Created = "20190621_074002"

	migration.Register("Blockchain_20190621_074002", m)
}

// Run the migrations
func (m *Blockchain_20190621_074002) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE blockchain (id SERIAL PRIMARY KEY, block bigint, index bigint);")
	m.SQL("INSERT INTO blockchain (block, index) VALUES (11681266, 0);")
}

// Reverse the migrations
func (m *Blockchain_20190621_074002) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS blockchain;")
}
