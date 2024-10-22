package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	const pi = math.Pi

	fmt.Print("Masukkan jumlah kerucut (n): ")
	fmt.Scan(&n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			var r, t float64
			fmt.Printf("Masukkan jari-jari alas dan tinggi untuk kerucut ke-%d (pisahkan dengan spasi): ", i)
			fmt.Scan(&r, &t)

			volume := (1.0 / 3.0) * pi * math.Pow(r, 2) * t

			fmt.Printf("Volume kerucut ke-%d adalah: %.2f\n", i, volume)
		}
	} else {
		fmt.Println("Masukkan jumlah kerucut yang valid (n harus lebih dari 0)!")
	}
}
