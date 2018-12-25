package main

import "fmt"

func main() {
	//fmt.Println("vim-go")
	//two 4 digit numbers:
	first := 1000
	last := 9999

	for init := first; init <= last; init++ {
		for i := init; i <= last; i++ {
			fmt.Println(i - init)
			//num := i * init
			//fmt.Println(num)
			//sprintf num ; if first and last chars match it is even ended number;
		}
	}
}
