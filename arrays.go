package main

import "fmt"

func arrays() {

	var arr1 [2]int
	var arr2 [3]int
	fmt.Println(arr1) // [0 0]
	fmt.Println(arr2) // [0 0 0]

	var arr3 [2]bool
	fmt.Println(arr3) // [false false]

	arr4 := [2]int{3, 9}
	fmt.Println(arr4) // [3 9]

	arr5 := [2][2]int{{3, 9}, {2, 4}}
	fmt.Println(arr5) // [[3 9] [2 4]]

	arr6 := [...][2]int{{1, 1}, {3, 9}, {2, 4}}
	fmt.Println(arr6)         // [[1 1] [3 9] [2 4]]
	fmt.Printf("%T \n", arr6) // [2][2]int

	//When arrays are passed into the functions Go sends a copy of the arrays to the function
}
