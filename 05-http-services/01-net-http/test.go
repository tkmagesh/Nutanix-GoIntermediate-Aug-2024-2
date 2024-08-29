package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	Id   int     `dbField:"product_id"`
	Name string  `dbField:"product_name"`
	Cost float64 `dbField:"product_cost"`
}

func main() {
	var p = Product{101, "Pen", 10}

	t := reflect.TypeOf(p)

	id_field := t.Field(0)

	id_tag := id_field.Tag.Get("dbField")
	fmt.Println(id_field.Name, id_tag)
}
