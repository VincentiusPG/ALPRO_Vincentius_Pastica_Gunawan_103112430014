package main

import "Fmt"

func main (){
	var input string
	fmt.Scan(&input)

	if input < "a" || input > "z" && input < "A" || input > "Z"{
		fmt.Print("Bukan huruf")
	} else if input == "A" || input == "a" || input == "I" || input == "i" || input == "U" || input == "u" || input == "E" || input == "e" || input == "O" || input == "o" {
		fmt.Print("Vokal")
	} else {
		fmt.Print("Konsonan")
	}	
}