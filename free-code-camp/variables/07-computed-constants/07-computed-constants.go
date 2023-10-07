package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	// Constants must be known at compile time.
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)
}
