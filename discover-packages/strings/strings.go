package main

import (
	"fmt"
	"strings"
	"unicode"
	"unsafe"
)

func main() {
	// func Clone
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)                                       // true
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone)) // false

	// func Compare
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1

	// func Contains
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true

	// func ContainsAny
	fmt.Println(strings.ContainsAny("team", "i"))     // false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     // true
	fmt.Println(strings.ContainsAny("failure", "ui")) // true
	fmt.Println(strings.ContainsAny("foo", ""))       // false
	fmt.Println(strings.ContainsAny("", ""))          // false

	// func ContainsFunc
	fmt.Println(strings.ContainsFunc("hello", func(r rune) bool { // true
		return string(r) == "h"
	}))

	// func ContainsRune
	fmt.Println(strings.ContainsRune("aardvark", 97)) // true
	fmt.Println(strings.ContainsRune("timeout", 97))  // false

	// func Count
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // 5, before & after each rune

	// func Cut
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")

	// func IndexRune
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1

	// func Join
	fmt.Println(strings.Join([]string{"foo", "bar", "baz"}, ", ")) // foo, bar, baz

	// func LastIndex
	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1

	// func LastIndexByte
	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) // 10
	fmt.Println(strings.LastIndexByte("Hello, world", 'o')) // 8
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) // -1

	// func LastIndexFunc
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) // 5
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) // 2
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))     // -1

	// func Map
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher...")) // 'Gjnf oevyyvt naq gur fyvgul tbcure...

	// func Repeat
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana

	// func Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo

	// func ReplaceAll
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) // moo moo moo

	// func Split
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	// func SplitAfter
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]

	// func SplitAfterN
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]

	// func SplitN
	fmt.Println(strings.SplitN("a,b,c", ",", 2))  // [a b,c]
	fmt.Println(strings.SplitN("a,b,c", ",", 0))  // []
	fmt.Println(strings.SplitN("a,b,c", ",", -1)) // [a b c]

	// func ToLower
	fmt.Println(strings.ToLower("Gopher")) // gopher

	// func ToLowerSpecial
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) // önnek iş

	// func ToUpper
	fmt.Println(strings.ToUpper("Gopher")) // GOPHER

	// func ToUpperSpecial
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş")) // ÖRNEK İŞ

	// func Trim
	fmt.Println(strings.Trim("!Hello World!", "!")) // Hello World

	// func TrimFunc
	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool { // Hello, Gophers
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// func TrimLeft
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers!!!

	// func TrimLeftFunc
	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool { // Hello, Gophers!!!
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// func TrimPrefix
	s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Println(s) // Gophers!!!

	// func TrimRight
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) // ¡¡¡Hello, Gophers

	// func TrimRightFunc
	fmt.Println(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool { // ¡¡¡Hello, Gophers
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// func TrimSpace
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers

	// func TrimSuffix
	s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Println(s) // ¡¡¡Hello
}
