package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	  
	// the orginal value is unchanged
	fmt.Println("orginal string value", greeting) 
	 
	// sort
	ages := []int{30, 50, 22, 10, 40}
	
	sort.Ints(ages)

	fmt.Println(ages)

	// read more -> https://golang.org/pkg/
}