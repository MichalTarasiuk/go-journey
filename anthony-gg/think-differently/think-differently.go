package main

import (
	"fmt"
	"time"
)

type Book struct {
	title    string
	author   string
	numPages int
	isSaved  bool
	savedAt  time.Time
}

func (book *Book) saveBooK() {
	book.isSaved = true
	book.savedAt = time.Now()
}

func main() {
	book := Book{
		title:    "Sample Title",
		author:   "Sample Author",
		numPages: 200,
	}

	book.saveBooK()

	fmt.Println(book)
}
