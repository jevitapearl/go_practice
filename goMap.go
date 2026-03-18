package main

import "fmt"

func goMap() {
	//Similar to dict in python

	mp1 := map[string]int{"a": 1}
	fmt.Println(mp1)

	mp2 := make(map[string]int)
	fmt.Println(mp2)

	mp3 := map[string][]int{"a": {1,2,3}} //maps having slices
	fmt.Println(mp3)

	mp4 := map[string][]int{"a": {1,2,3,4}}
	mp4["b"] = []int{3,4,5,6}
	fmt.Println(mp4)
	fmt.Println(mp4["b"]) //[3 4 5 6]

	delete(mp4, "a")
	fmt.Println(mp4)

	value, ok  := mp4["a"]
	fmt.Println(value,ok)

	//building a func where you check how many values from 1-100 is divisible by the key
	mp5 := map[uint]uint{}
	n:= 100

	for i:=1; i<=n ; i++{
		for k := 1; k <= 20; k+=2{
			if i%k == 0{
				mp5[uint(k)]++
			}
		}
	}
	fmt.Println(mp5)

}
