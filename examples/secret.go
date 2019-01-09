// +build ignore

package main

//go:generate strobfus -filename $GOFILE

import "fmt"

// a little comment
var hello string

var yolo = "poeut"

var helloWorld = "hello\nworld"

var arr = []string{
	"a",
	"b",
}

var (
	str1 = "coucou"
	arr1 = []string{
		"foo",
		"bar",
	}
	arr2 = []string{"h", "g"}
	arr3 = []string{"unique entry"}
)

func init() {
	fmt.Println("This my init")
}
