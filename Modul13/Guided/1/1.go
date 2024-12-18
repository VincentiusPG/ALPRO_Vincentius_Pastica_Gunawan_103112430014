package main

import "fmt"

func main () {
	var input int
	fmt.Print("masukan input")
	fmt.Scan(&input)

	for i := 1 ; i <= input ; i++ {
		if i % 2 != 0 {
			fmt.Println(i)
		}
	}
}