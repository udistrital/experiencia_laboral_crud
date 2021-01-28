package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsParametricas_20190804_190436 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsParametricas_20190804_190436{}
	m.Created = "20190804_190436"

	migration.Register("InsertsParametricas_20190804_190436", m)
}

// Run the migrations
func (m *InsertsParametricas_20190804_190436) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190804_190436_inserts_parametricas.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *InsertsParametricas_20190804_190436) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190804_190436_inserts_parametricas.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
