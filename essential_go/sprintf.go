
package main

import (
	"fmt"
)

func main() {
	n := 42
	fmt.Printf("n = %q (type %T) \n", n, n)
	s := fmt.Sprintf("%d", n)
	fmt.Printf("s = %q (type %T) \n", s, s)
}
