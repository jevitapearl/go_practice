package main

import "fmt"

func variables(){
	//uint - 8, 16, 32, 64 - default(32)
	//int - 8, 16, 32, 64 - default(32)
	//float - 8, 16, 32, 64 - default(32)
	//byte = int8 = 1byte
	//rune = int32
	//bool - true, false
	//string - "" only not ''
	//nil - null in JS/None in Python

	//Explicit assignment
	//var/const varName dtype
	var x string = "Hello. This is where I'm learning about variables and datatypes"
	const y int32 = 470
	fmt.Println(x, y)

	//Implicit assignment
	z := 50
	fmt.Println("This is an implicit variable value", z)
	fmt.Printf("Type %T\n", z)
	//%T - type
	//%v - value
	//%b - binary rep
	//%e - scientific notation
	//%f - float
	//%s - string
	//%.4f - float with 4 precision
	//%% - % in a string

	storedStr := fmt.Sprintf("Value of x = \"%v\"", x) //Stores the formatted string
	fmt.Println(storedStr)

	//TYPECASTING
	//typecasting is same as python
	//Arithmetic operations should be done by converting all the operands into a single type

	//e.g.: uint8  + int32 => error
	//int/int => int
	//float64/float64 => float64

	// Stirng type-casting

	a := "hi"
	b := 3
	c := a + string(b)
	fmt.Println(c)
	//hi

	p := "hi"
	q := 3
	r := p + fmt.Sprint(q)
	fmt.Println(r)
	//hi3

	//Comparison between variables is only possible if the types are same
	//e.g.: int8 <= int 32 ----> error
}
