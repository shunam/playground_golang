package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20171022_145316 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20171022_145316{}
	m.Created = "20171022_145316"

	migration.Register("Users_20171022_145316", m)
}

// Run the migrations
func (m *Users_20171022_145316) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) NOT NULL AUTO_INCREMENT,`email` varchar(128) NOT NULL,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Users_20171022_145316) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
