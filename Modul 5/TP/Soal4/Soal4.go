package main

import (
	"fmt"
)

func main() {

	randomNumber := 69

	var guess int
	const maxAttempts = 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih angka antara 1 hingga 100.")
	fmt.Printf("Anda memiliki %d kali kesempatan untuk menebak.\n", maxAttempts)

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Tebakan %d: ", attempts)
		fmt.Scan(&guess)

		if guess == randomNumber {
			fmt.Println("Selamat! Tebakan Anda benar!")
			break
		} else if guess < randomNumber {
			fmt.Println("Angka terlalu kecil.")
		} else {
			fmt.Println("Angka terlalu besar.")
		}
		if attempts == maxAttempts {
			fmt.Printf("Kesempatan habis! Angka yang benar adalah %d.\n", randomNumber)
		}
	}
}
