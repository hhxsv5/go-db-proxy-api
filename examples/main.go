package main

import (
	"github.com/hhxsv5/go-db-proxy-api/drivers"
	"fmt"
)

func main() {
	m := drivers.NewMySQL("user:password@tcp(120.0.0.1:3306)/dbname")
	err := m.Open()
	if err != nil {
		panic(err.Error())
	}
	defer m.Close()

	rows := m.Query("select id,cellphone from user order by id desc limit 5")

	fmt.Println(rows)

}
