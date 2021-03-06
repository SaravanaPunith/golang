package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1,2,3,4}
	doubleAt(values, 2)
	fmt.Println(values)
	
	value:=4
	double(value)
	fmt.Println(value)

	doublePtr(&value)
	fmt.Println(value)
}
