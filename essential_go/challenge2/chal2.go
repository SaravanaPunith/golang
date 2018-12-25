package main

import "fmt"

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("n = %d type %T \n", n, n)
	fmt.Printf("s = %s type %T \n", s, s)
	fmt.Printf("s = %q type %T \n", s, s)
}
