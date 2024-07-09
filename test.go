package main

import (
	"fmt"
	"project/mypackage"
	"unicode/utf8"
)

var a int8 = 1

func main() {
	fmt.Println(utf8.RuneCountInString("123\""))
	var a int8 = 1
	//a := 1
	fmt.Println(a)
	// 调用 MyFunction
	mypackage.MyFunction()
	mypackage.NameReturn("Bob", 12)
}
