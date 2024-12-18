package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	// Input tiga bilangan dari pengguna
	fmt.Println("Masukkan tiga bilangan:")
	fmt.Print("Bilangan pertama: ")
	fmt.Scan(&num1)
	fmt.Print("Bilangan kedua: ")
	fmt.Scan(&num2)
	fmt.Print("Bilangan ketiga: ")
	fmt.Scan(&num3)

	// Menentukan nilai terbesar dan terkecil
	max, min := num1, num1

	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}

	if num2 < min {
		min = num2
	}
	if num3 < min {
		min = num3
	}

	// Menampilkan hasil
	fmt.Print("Nilai terbesar adalah: ", max)
	fmt.Print("Nilai terkecil adalah: ", min)
}
