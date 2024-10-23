package main

import "fmt"

func main(){
	var input, dinar, dirham, fals,qirat int
	fmt.Println("masukan qirat")
	fmt.Scan(&input)
	dinar = input / 600
	dirham = (input % 600) / 60
	fals = (input % 60) / 6
	qirat = input % 6
	fmt.Println(dinar,dirham,fals,qirat)
}