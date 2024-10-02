package main

import "fmt"

func main(){
	var input int
	var angka1 float32
	var angka2 float32

	fmt.Println("1.Penjumlahan")
	fmt.Println("2.Pengurangan")
	fmt.Println("3.Perkalian")
	fmt.Println("4.Pembagian")
	fmt.Println("Silahkan pilih opsi aritmatika yang ingin digunakan:")
	fmt.Scan(&input)

	if input == 1 {
		fmt.Println("Masukan angka pertama:")
		fmt.Scan(&angka1)
		fmt.Println("Masukan angka kedua:")
		fmt.Scan(&angka2)
		fmt.Println("Hasil penjumlahannya adalah", angka1 + angka2)
	} else if input == 2 {
		fmt.Println("Masukan angka pertama:")
		fmt.Scan(&angka1)
		fmt.Println("Masukan angka kedua:")
		fmt.Scan(&angka2)
		fmt.Println("Hasil pengurangannya adalah", angka1 - angka2)
	} else if input == 3 {
		fmt.Println("Masukan angka pertama:")
		fmt.Scan(&angka1)
		fmt.Println("Masukan angka kedua:")
		fmt.Scan(&angka2)
		fmt.Println("Hasil perkaliannya adalah", angka1 * angka2)
	} else if input == 4 {
		fmt.Println("Masukan angka pertama:")
		fmt.Scan(&angka1)
		fmt.Println("Masukan angka kedua:")
		fmt.Scan(&angka2)
		fmt.Println("Hasil pembagiannya adalah", angka1 / angka2)
	} 
}