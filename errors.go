package main

import "fmt"

func deferredFunc1() {
	fmt.Println("defer 1")
	if r := recover(); r != nil {
		fmt.Printf("Panicked and recovered: %v\n", r)
	}
}

func deferredFunc2() {
	fmt.Println("defer 2")
	if r := recover(); r != nil {
		fmt.Printf("Panicked and recovered: %v\n", r)
	}
}

func deferredFunc3() {
	fmt.Println("defer 3")
	if r := recover(); r != nil {
		fmt.Printf("Panicked and recovered: %v\n", r)
	}
}

func panicNow() {
	panic("This caused a crash")
}

func errors() {
	defer deferredFunc1() //acts as finally before crashing the function
	defer deferredFunc2()
	defer deferredFunc3() 
	// breaks the flow of the program. Recovery code should be before the panic
	panicNow()

	//defer id stacked and recover takes place in the same way
}
