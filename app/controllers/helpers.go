package controllers

import (
	"fmt"
	"reflect"
)

// Mapable .
type Mapable interface{}

func extract(attrs []string, list ...interface{}) {
	fmt.Println(reflect.ValueOf(list))
}
