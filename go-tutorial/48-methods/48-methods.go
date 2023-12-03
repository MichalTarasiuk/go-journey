package main

import "fmt"

type User struct {
	Name string
}

func (user User) sayHello() {
	fmt.Printf("hello %s\n", user.Name)
}

func main() {
	user := User{
		Name: "John",
	}

	user.sayHello()
}
