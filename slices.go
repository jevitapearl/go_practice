package main

import "fmt"

func slices() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	s1 := arr1[:3]
	fmt.Println(s1) // [1 2 3]

	fmt.Println(arr1[2:]) // [3 4 5]
	fmt.Println(arr1[:])  // [1 2 3 4 5]

	s1[0] = 10
	fmt.Println(arr1) // [10 2 3 4 5]

	// pointer
	fmt.Println(s1)

	// length
	fmt.Println(len(s1))

	//capacity
	fmt.Println(cap(s1))

	// Slices are the view of the underlying array

	s2 := []string{"hello", "world"}

	for x := 0; x < 10; x++ {
		s2 = append(s2, "what")
		fmt.Println(s2, len(s2), cap(s2)) // capacity always doubles as elements are added
	}

	s3 := make([]int, 10)
	fmt.Println(s3) // [0 0 0 0 0 0 0 0 0 0]

	s4 := []string{"hello", "you"}
	fmt.Println(s4) // [hello you]
	test(s4)
	fmt.Println(s4) // [changed this you]

	// When slices are passed as args the value of the array is changed if it is messed with in the function

}

func test(arr []string) {
	arr[0] = "changed"
}
