package main

import (
	"fmt"
	"math"
)

func main() {
	var x = 1 // int32 or int64 depending on machine
	var x32 int32 = 1
	var x64 int64 = 1

	fmt.Println(x, x32, x64)
	fmt.Printf("%T, %T, %T", x, x32, x64)

	x = int(x32)
	fmt.Println(x)
	x = int(x64)
	fmt.Println(x)

	var unsigned uint = 1
	// error: unsigned  = -1
	fmt.Println(unsigned)

	y := 2
	fmt.Println(x + y)  // 3
	fmt.Println(x % y)  // 1
	fmt.Println(x == y) // fasle
	fmt.Println(x < y)  // true

	// floating point numbers
	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi)
	fmt.Println()

	var pi32 float32 = 3.141
	fmt.Printf("%T", pi32)

	pi = float64(pi32)
	fmt.Println(pi)

	// conversion between the two
	a := 1
	b := 3.999
	fmt.Println(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	// math library
	fmt.Println(math.Min(0, 1))      // 0
	fmt.Println(math.Max(0, 1))      // 1
	fmt.Println(math.Ceil(3.00001))  // 4
	fmt.Println(math.Floor(3.00001)) // 3
	fmt.Println(math.Abs(-1))        // 1
	fmt.Println(math.Sqrt(100))      // 10
	fmt.Println(math.Pow(2, 3))      // 8
}
