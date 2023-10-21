package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello World"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Printf("%c", s[0])
	fmt.Println()
	fmt.Println(s[0:6])
	fmt.Println(s[:6])
	fmt.Println(s[6:])

	s = s + " Again"
	s += " Again"
	fmt.Println(s)

	fmt.Println("Hello \nWorld")
	fmt.Println("Hello \tWorld")
	fmt.Println("Hello \bWorld")

	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s, %s", abc, b)
	fmt.Println()
	abc2 := string(b)
	fmt.Println(abc2)

	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))
	fmt.Println(strings.Replace("Hello World World", "World", "Hello", 1))
	fmt.Println(strings.Replace("Hello World World", "World", "Hello", 2))

	var sb strings.Builder
	fmt.Println("This is a String Builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("Hello")
	fmt.Println("This is a String Builder:", sb.String())
	fmt.Println(sb.Cap()) // 0 2 4 8
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println(sb.String())

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println()
	z, _ := strconv.Atoi(y)
	fmt.Printf("%T", z)
}
