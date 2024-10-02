package main

import "fmt"

func main (){
	var sisi int
	var hasil int
	fmt.Println("masukan sisi kubus:")
	fmt.Scan(&sisi)
	hasil = sisi * sisi * sisi
	fmt.Println("Voleme nya adalah : ", hasil)
}