package main

import "fmt"

func conditionalsAndLoops() {
	a := 1
	switch a {
	case 1:
		fmt.Println("The value is 1")
		fallthrough //slips the control into the next case
	case 2:
		fmt.Println("The value is 2")
		fallthrough
	case 3:
		fmt.Println("The value is 3")
	}

	//switch with no follow-up variable will act like a if-elif-else ladder

	//Strings have 2 formats: ASCII-1byte and UTF-8(4bytes - for chars including emojis)

	s := "hello to you sir"
	fmt.Println(s[2])         //prints the ASCII value of the character
	fmt.Println(s)            //prints the string
	fmt.Println(string(s[2])) //prints the character

	for i := range s {
		fmt.Printf("%c", s[i]) //prints the character one-by-one by index
	}
	fmt.Println()

	for i, _ := range s {
		fmt.Printf("%c", s[i]) //We're using the index and not using the characters directly
	}
	fmt.Println()

	for _, c := range s {
		fmt.Printf("%c", c) //We're using the characters directly and not index
	}
	fmt.Println()

	for idx := 0; idx < len(s); idx++ {
		fmt.Printf("%c", s[idx]) //Acts like a traditional for loop
	}
	//fmt.Println(idx) //Gives an error because the idx was only existing inside the loop
	fmt.Println()

	idx := 0
	for idx < len(s) {
		fmt.Printf("%c", s[idx]) //prints the character one-by-one
		idx++
	}

	fmt.Println()
	fmt.Println(idx) //len of the string

}
