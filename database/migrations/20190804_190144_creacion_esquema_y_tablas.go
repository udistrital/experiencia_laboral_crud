package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionEsquemaYTablas_20190804_190144 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionEsquemaYTablas_20190804_190144{}
	m.Created = "20190804_190144"

	migration.Register("CreacionEsquemaYTablas_20190804_190144", m)
}

// Run the migrations
func (m *CreacionEsquemaYTablas_20190804_190144) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190804_190144_creacion_esquema_y_tablas.up.sql")

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
func (m *CreacionEsquemaYTablas_20190804_190144) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190804_190144_creacion_esquema_y_tablas.down.sql")

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
