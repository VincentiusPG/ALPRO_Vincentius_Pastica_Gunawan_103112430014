package main

import "fmt"

func main () {
	var kata int

	for done := false; !done; {
		fmt.Scan(&kata)
		done = (kata > 0)
	} 
	fmt.Println (kata , "merupakan bilangan positif")
}