package main

import (
	"fmt"
	"reflect"
)

type anderson int

var x anderson

func main() {
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	x = 42
	fmt.Println(x)

}
