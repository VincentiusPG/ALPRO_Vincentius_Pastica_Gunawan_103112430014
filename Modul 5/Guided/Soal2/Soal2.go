package main 

import "fmt"

func main(){
	var alas, tinggi, n, i int
	var segitiga float32
	fmt.Print("Masukan banyaknya perulangan")
	fmt.Scan(&n)

	for i = 1 ; i <= n ; i++{
		fmt.Scan(&alas, &tinggi)
		segitiga = 0.5 * float32(alas*tinggi)
		fmt.Println(segitiga)
	}
}