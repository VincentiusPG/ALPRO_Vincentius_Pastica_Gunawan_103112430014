package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan angka 4 digit: ")
	fmt.Scan(&input)

	if input >= 1000 && input <= 9999 {
		a1 := input / 1000         
		a2 := (input / 100) % 10   
		a3 := (input / 10) % 10    
		a4 := input % 10           

		if a1 > a2 && a2 > a3 && a3 > a4 {
			fmt.Println("Menurun")
		} else if a1 < a2 && a2 < a3 && a3 < a4 {
			fmt.Println("Membesar")
		} else {
			fmt.Println("Tidak terurut")
		}
	} else {
		fmt.Println("Input harus 4 digit")
	}
}