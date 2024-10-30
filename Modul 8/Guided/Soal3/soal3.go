package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	if bilangan % 2 == 0 && bilangan < 0 {
		fmt.Println("Bilangan tersebut adalah bilangan negatif genap.")
	} else {
		fmt.Println("Bilangan tersebut bukan bilangan negatif genap.")
	}
}
