package main

import (
	"fmt"
)

func main() {
	var number int

	// Input bilangan dari pengguna
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	// Menampilkan faktor-faktor bilangan
	fmt.Printf("Faktor-faktor dari %d adalah:\n", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}