// slice - sequence of items
// all items in slice same type

// append - add to slice
// slices on top of array.
//  

package main


import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons %v type %T\n", loons, loons)

	fmt.Println(len(loons))

	// for
	for i:=0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}
	
	for j:= range loons {
		fmt.Println(j)
	}

	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	fmt.Println("********")
	for _, name := range loons {
		fmt.Println(name)
	}
	fmt.Println("-----")

	fmt.Println(loons[1])

	fmt.Println("=======")
	fmt.Println(loons[1:])

	loons = append(loons, "newvalue")
	fmt.Println(loons)
}

