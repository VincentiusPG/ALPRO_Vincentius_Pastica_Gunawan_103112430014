package main 

import "fmt"

func main(){
	var x,hasil  int
	fmt.Println("Masukan nilai x:")
	fmt.Scan(&x)
	for i := 1 ; i <= x ; i+=1 {
		hasil = i * i
		fmt.Println(hasil)
	}
}