package main

import "fmt"

func main() {
	var base, exponent int
	fmt.Print("Masukkan bilangan dasar (base): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan bilangan pangkat (exponent): ")
	fmt.Scan(&exponent)

	if base >= 0 && exponent >= 0 {
		result := 1
		for i := 0; i < exponent; i++ {
			result *= base
		}
		fmt.Printf("%d dipangkatkan %d adalah: %d\n", base, exponent, result)
	} else {
		fmt.Println("Masukkan bilangan bulat positif untuk base dan exponent!")
	}
}
