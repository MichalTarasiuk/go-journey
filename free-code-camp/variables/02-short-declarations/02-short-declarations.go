package main

import "fmt"

// / Outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func main() {
	congrats := "happy birthday!"
	fmt.Println(congrats)
}
