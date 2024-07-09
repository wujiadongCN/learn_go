package mypackage

import "fmt"

func MyFunction() {
	fmt.Println("This is a function in mypackage.")
}

func NameReturn(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}
