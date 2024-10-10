package main

import (
	"fmt"
	"math"
)

func main() {
	var xA, yA, xB, yB, xC, yC float64
	
	fmt.Println("Masukkan koordinat titik A (x y):")
	fmt.Scan(&xA, &yA)
	
	fmt.Println("Masukkan koordinat titik B (x y):")
	fmt.Scan(&xB, &yB)
	
	fmt.Println("Masukkan koordinat titik C (x y):")
	fmt.Scan(&xC, &yC)
	
	AB := math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	BC := math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	AC := math.Sqrt(math.Pow(xC-xA, 2) + math.Pow(yC-yA, 2))
	
	longest := AB
	if BC > longest {
		longest = BC
	}
	if AC > longest {
		longest = AC
	}
	
	fmt.Printf("Panjang sisi terpanjang adalah %.2f", longest)
}
