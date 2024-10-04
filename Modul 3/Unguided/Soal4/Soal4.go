package main

import "fmt"

func main(){
	var C float32
	var R float32
	var F float32
	var K float32
	fmt.Println("Masukan nilai Celcius: ")
	fmt.Scan(&C)
	R = 0.8 * C
	F = 1.8 * C + 32
	K = C + 273
	fmt.Println("Temperatur Celcius : ",C)
	fmt.Println("Derajat Reamur : ",R)
	fmt.Println("Derajat Fahrenheit : ",F)
	fmt.Println("Derajat Kelvin : ",K)
}