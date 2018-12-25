package main

import (
	"fmt"
)

func main() {
	// int8 int16 int32 int64  and unsigned
	// int depends on system you are using

	// go type system is strict.
	/*
	   var x float64
	   var y float64
	*/
	// zero by default
	/*
	   x := 1.0
	   y := 2.0 */
	x, y := 1.0, 2.0
	/*
	   // z declared not used
	   z:= 10
	*/
	// %v value of object %T type
	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	// 1+2 /2 =1 and not 1.5 as expected
	mean := (x + y) / 2
	fmt.Printf("result = %v, type of %T\n", mean, mean)
}
