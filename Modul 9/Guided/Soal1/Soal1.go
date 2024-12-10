package main

import "fmt"

func main() {
	var umur int
	var kk bool

	fmt.Scan(&umur)
	fmt.Scan(&kk)

	if umur >= 17 && kk == true {
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("Tidak bisa membuat ktp")
	}
}