package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	if N > 0 {
		for i := 1; i <= N; i++ {
			fmt.Printf("Kuadrat dari %d adalah %d\n", i, i*i)
		}
	} else {
		fmt.Println("Harap masukkan bilangan bulat positif!")
	}
}