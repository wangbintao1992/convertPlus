package main

import (
	"fmt"
	"testing"
)

type User struct {
	Name string `1convert:cached=true,order=0,defaultValue:{Age:123},doubleFormat:0.00`
	Age  int    `json:"age" bson:"b_age"`
}

// aTest
func TestBase(t *testing.T) {
	//c = Convert
	//Convert.scanSourceField()
	fmt.Println("2121")
}
