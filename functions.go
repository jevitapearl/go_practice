package main

import "fmt"

func callFunc(callable func(int) int) int { //passing func as a parameter
	return callable(10) //sends 10 as an arguement to the function that was passed to this
}

func doubleNumber(x int) int {
	return x * 2
}

func getFunc(s1 string) func(string) string {
	return func(s2 string) string {
		return s1 + s2
	}
}

func sum(nums ...int) (int, int) {
	s := 0
	for _, value := range nums {
		s += value
	}
	return s, s * s
}

func functions() {
	value1 := callFunc(doubleNumber) //sending the func doubleNumber to the callable func
	fmt.Println(value1)

	f1 := getFunc("hello ") //returns a function which can be executed by calling f1
	value2 := f1("world")   // the returned function adds the value previously passed while calling f1 and concatenates with the current value of arguement
	fmt.Println(value2)

	value3, value4 := sum([]int{1, 2, 3, 4, 5}...)
	fmt.Println(value3, value4)
}
