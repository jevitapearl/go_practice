package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to this to-do app")
	http.HandleFunc("/", helloUser)
	http.ListenAndServe(":8080", nil) 
}

func helloUser(w http.ResponseWriter, req *http.Request) {
	var greeting = "hello to you user"
	var taskItems = []string{"Push to git", "Build a basic System design", "Take rewards"}


	fmt.Fprintln(w, greeting)
	for index, task := range taskItems {
		fmt.Fprintf(w, "%d. %s\n", index+1, task)
	}
}
