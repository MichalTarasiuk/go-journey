package main

import "fmt"

// When
// 1. When we need to update state
// 2. when we want to optimize the memory for Large objects that are getting called A LOT.

type User struct {
	email string
}

func (user *User) updateEmail(newEmail string) {
	user.email = newEmail
}

func main() {
	user := User{
		email: "example@gmail.com",
	}

	user.updateEmail("example2@gmail.com")

	fmt.Println(user)
}
