package main

import "fmt"

func main (){
	var fahrenheit, kelvin int
	fmt.Println("Masukan suhu dalam derajat fahrenheit:")
	fmt.Scan(&fahrenheit)
	kelvin = 5/9 * (fahrenheit-32) + 273
	fmt.Println("Konversi ke kelvin dari ",fahrenheit,"fahrenheit, menjadi ", kelvin,"kelvin")
}

