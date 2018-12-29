
package main

import (
	"fmt"
)

func main() {

	stocks := map[string]float64 {
		"AMZN" : 169.8,
		"GOOG" : 1129,
		"MSFT" : 98.1,
	}
	fmt.Println(len(stocks))

	fmt.Println(stocks["MSFT"])
	fmt.Println(stocks["TSLA"])

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stocks["TSLA"] = 322.12
	value, ok = stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	delete (stocks,"AMZN")
	fmt.Println(stocks)

	for key:= range stocks {
		fmt.Println(key)
	}
	
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f", key, value)
	}
}

// map is a ds key, with value
// keys must be of same type

//value of same type

