package main

import "fmt"

func main() {
	book := "The color of magic"

	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T) \n", book, book)
	fmt.Printf("book[0] = %v (type %T) \n", book[0], book[0])

	// strings immutable - once assigned cannot be changed
	// book[0] = 116
	// not including 11
	fmt.Println(book[4:11])
	// from 4 till end
	fmt.Println(book[4:])
	// from begin .. not including 4
	fmt.Println(book[:4])

	fmt.Println("t" + book[1:])

	fmt.Println("It was 1/2 price")

	poem := ` the road goes on
down from the door
where it begins...
		 `
	fmt.Println(poem)
}
