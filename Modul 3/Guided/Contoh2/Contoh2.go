package main

import "fmt"

func main (){
	var alas float32
	var tinggi float32
	var hasil float32

	fmt.Println("masukan alas:")
	fmt.Scan(&alas)
	fmt.Println("masukan tinggi:")
	fmt.Scan(&tinggi)

	hasil= 0.5 * alas * tinggi
	fmt.Println("Luas segitiga adalah", hasil)
}