package main

import (
	"fmt"
	"reflect"
)

func main() {
	//scanner := NewScanner("1 + 3")
	//tokens := scanner.Tokens
	//errors := scanner.Errors
	//done := scanner.Done
	//go scanner.Scan()
	//
	//<- done
	//fmt.Println("Tokens:")
	//for t := range tokens {
	//	fmt.Println(t)
	//}
	//fmt.Println("Errors:")
	//for e := range errors {
	//	fmt.Println(e)
	//}

	fmt.Println(reflect.TypeOf("hi"[0]))
	fmt.Println(reflect.TypeOf('h'))
}
