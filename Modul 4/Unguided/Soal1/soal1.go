package main

import "fmt"

func main(){
	var belanja, diskon, potongan, hasil float32
	fmt.Println("Masukan total belanja dan ptotngan diskon:")
	fmt.Scan(&belanja,&diskon)
	potongan = belanja * (diskon/100)
	hasil = belanja - potongan
	fmt.Println("Total akhir adalah ", hasil)
}