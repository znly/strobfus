// +build ignore

package main

//go:generate strobfus -filename $GOFILE

// To generate this file deterministically, add the option '-seed foobar',
// 'foobar' being a seed from which the AES keys will derive.

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
