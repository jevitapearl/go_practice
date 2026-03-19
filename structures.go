package main

import "fmt"

type Person1 struct {
	name string
	age  uint
}

type Person2 struct {
	name string
	age  uint
	f    func(string) string
}

func getName(p Person1) string { //accepts a struct as an args and returns a value associated with it
	return p.name
}

func (p Person1) getName() string {
	return p.name
}

func (p Person1) setName(newName string){
	p.name = newName
	fmt.Println(p.name)
}

func structures() {
	p1 := Person1{"jev", 20}
	fmt.Println(p1) // {jev 20}

	p1.name = "JP"
	fmt.Println(p1) //{JP 20}

	var p2 Person1 = Person1{age: 40}
	fmt.Println(p2) // { 40}
	p2.name = "Poirot"
	fmt.Println(p2) // {Poirot 40}

	fmt.Println(getName(p1))

	p3 := Person2{}
	p3.f = func(x string) string {
		return x + "s"
	}
	fmt.Println(p3.f("her")) // hers

	p4 := Person1{name: "pearl", age: 33}
	fmt.Println(getName(p4))  // regular function accepting a struct as an arg
	fmt.Println(p4.getName()) // object behaviour

	p4.setName("new jev") //send a copy of the struct and not the ref
	fmt.Println(p4.name) // p4.name is not modified and remains the same


	//anything starting with uppercase is exportable
	// similar to private and public variables and methods
}
