package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address *AddrDetail
}
type AddrDetail struct {
	Province string
	City     string
}

var people = Person{
	Name: "wafa",
}

func TestNormalSample(t *testing.T) {
	fmt.Printf("1, people: %#v\n", people)
	// AddAge(20, addr)
	AddAge(20, people.Address)
	fmt.Printf("2, people: %#v\n", people)
}
func AddAge(age int, addr *AddrDetail) {
	addr = &AddrDetail{City: "洛阳", Province: "河南"}
	// people.Address = addr
}
