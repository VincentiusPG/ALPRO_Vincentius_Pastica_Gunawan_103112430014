package main

import "fmt"

func main (){
	var Celcius , Fahrenheit int

	fmt.Println("Masukan suhu fahrenheit :")
	fmt.Scan(&Fahrenheit)
	Celcius = 5/9 * (Fahrenheit-32)

	fmt.Println(Fahrenheit, "Fahrenheit dalam celcius adalah :", Celcius,"Celcius")
}