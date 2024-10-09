package main

import "fmt"

func main (){
	var berat, tinggi, hasil float32
	fmt.Print("masukan berat badan dan tinggi badan:")
	fmt.Scan(&berat,&tinggi)
	hasil = berat / (tinggi * tinggi)
	fmt.Printf("BMI nya adalah %0.2f", hasil) // %0.2f digunakan agar hasil desimalnya hanya mengambil 2 digit dibelakang koma(angka 2 dapat dinganti dengan angka lain dan dapat disesuaikan kebutuhan)
}