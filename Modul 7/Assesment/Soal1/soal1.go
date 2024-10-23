package main

import "fmt"

func main(){
	var x, y, hasil int
	fmt.Println("Masukan nilai x dan y:")
	fmt.Scan(&x,&y)
	for i := x ; i <= y ; i+=1 {
		hasil = hasil + i
		fmt.Println(hasil)
	}
}