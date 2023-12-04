package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func getInstance() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func main() {
	s := getInstance()

	s["this"] = "that"

	s2 := getInstance()

	fmt.Println("This is", s2["this"]) // This is that
}
