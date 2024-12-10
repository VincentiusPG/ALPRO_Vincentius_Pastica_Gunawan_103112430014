package main

import "fmt"

func main() {
	var berat int        
	var totalBeratKg int  
	var sisaBeratGram int 
	var biayaKg int       
	var biayaSisa int     
	var totalBiaya int   

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	
	totalBeratKg = berat / 1000
	sisaBeratGram = berat % 1000

	
	biayaKg = totalBeratKg * 10000 

	
	if totalBeratKg > 10 {
	
		biayaSisa = 0
	} else {
		if sisaBeratGram >= 500 {
			biayaSisa = sisaBeratGram * 5 
		} else {
			biayaSisa = sisaBeratGram * 15 
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", totalBeratKg, sisaBeratGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n",totalBiaya)
}