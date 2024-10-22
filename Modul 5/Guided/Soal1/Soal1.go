package main

import "fmt"

func main(){
	var i, n ,input1 int
	fmt.Print("masukan angka awal dan akhir")
	fmt.Scan(&input1,&n)

	for i = input1 ; i <= n ; i++{
		fmt.Print(i)
	}
}