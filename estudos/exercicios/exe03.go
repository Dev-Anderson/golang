package main

import "fmt"

func main() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)

}
