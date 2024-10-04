package main

import "fmt"

func main() {
    var tahun int
    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&tahun)
    Kabisat := false
    if (tahun % 400 == 0) || (tahun % 4 == 0 && tahun % 100 != 0) {Kabisat = true}

    // Menampilkan hasil
    fmt.Println("Tahun : ", tahun)
	fmt.Println("Kabisat : ", Kabisat)
}