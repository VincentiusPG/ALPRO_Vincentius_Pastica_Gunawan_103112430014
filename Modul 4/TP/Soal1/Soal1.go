package main

import "fmt"

func main(){
	var r,luas,keliling float32
	fmt.Println("Masukan nilai jari jari:")
	fmt.Scan(&r)
	luas = 3.14 * r * r
	keliling = 2 * 3.14 * r
	fmt.Println("luas : ", luas, " dan keliling : ", keliling) 
}