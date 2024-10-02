package main

import "fmt"

func main(){
	var idr float32
	var usd float32
	fmt.Println("Masukan idr")
	fmt.Scan(&idr)
	usd = idr /15000
	fmt.Println(usd,"dollar")
}