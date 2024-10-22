package main

import "fmt"

func main() {
	var jumlahBarang int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)
	totalPoin := 0
	if jumlahBarang > 5 {
		totalPoin += 5 * 10
		totalPoin += (jumlahBarang - 5) * 15
	} else {
		totalPoin += jumlahBarang * 10
	}

	fmt.Printf("Total poin yang didapatkan: %d poin\n", totalPoin)
}