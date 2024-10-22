package main

import "fmt"

func main() {
	var n int
	var sum int

	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			sum += i
		}
		fmt.Printf("Penjumlahan dari 1 sampai %d adalah %d\n", n, sum)
	} else {
		fmt.Println("Masukkan bilangan bulat positif!")
	}
}
