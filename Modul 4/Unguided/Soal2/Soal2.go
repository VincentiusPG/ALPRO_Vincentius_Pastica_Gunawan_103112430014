package main

import "fmt"

func main() {
	var bb, bmi, tinggi float32
	fmt.Println("masukan nilai BMI dan tinggi badannya(meter)")
	fmt.Scan(&bmi, &tinggi)
	bb = bmi * (tinggi * tinggi)
	fmt.Println("Berat badan adalah", bb, "kg")
}