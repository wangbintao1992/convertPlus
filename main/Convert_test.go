package main

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string `convert:cached=true,order=0,defaultValue:{Age:123},doubleFormat:0.00`
	Age  int    `json:"age" bson:"b_age"`
}

// aTest
func TestBase(t *testing.T) {
	var u User
	var u2 User
	u.Name = "name"
	u.Age = 21
	t1 := reflect.TypeOf(u2)
	convert := new(Convert)
	//var t2 any
	convert.convert(u, t1)

	//c = Convert
	//Convert.scanSourceField()
	//Convert.New
	fmt.Println("2121")
}
