package main

import "fmt"

func main(){
	var a, b , hasil int
	fmt.Print("Masukan nilai a dan b") 
	fmt.Scan(&a, &b)
	for i	 := 1 ; i <= b ; i++{
		hasil = hasil + a
	}
	fmt.Print(hasil)
}
