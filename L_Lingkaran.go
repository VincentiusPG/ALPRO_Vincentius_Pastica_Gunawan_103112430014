package main

import "fmt"

func main () {
	var r float32
	var L float32

	fmt.Println("masukan jari jari lingkaran :")
	fmt.Scan(&r)

	L = 3.14 * r * r
	fmt.Println("Hasil luas lingkaran dari jari jari :", r , "adalah", L)
}