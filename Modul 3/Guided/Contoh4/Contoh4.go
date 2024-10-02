package main

import "fmt"

func main(){
	var x int
	var fungsi int
	fmt.Println("Masukan nilai x")
	fmt.Scan(&x)
	fungsi = 2/(x+5) + 5
	fmt.Println("Hasil dari fungsi x adalah", fungsi) 
}