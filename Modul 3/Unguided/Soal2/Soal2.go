package main

import "fmt"

func main(){
	var r float32
	var vol float32
	var lbola float32

	fmt.Print("masukan nilai jari jari bola:")
	fmt.Scan(&r)
	vol = (4.0/3.0) * 3.1415926535 * r * r * r
	lbola = 4.0 * 3.1415926535 * r * r
	fmt.Print("Jejari = ", r)
	fmt.Print("Bola denngan jejari ", r, " memiliki volume ", vol, " dan luas kulit ", lbola) 
}