package main

import "fmt"

func main(){
	var x float32
	var fx float32
	fmt.Print("Masukan nilai fx:")
	fmt.Scan(&fx)
	x = (2 / (fx - 5)) - 5
	fmt.Print("x adalah :", x)
}