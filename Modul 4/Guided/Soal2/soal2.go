package main 

import "fmt"

func main (){
	var digit1,digit2,digit3 int
	var input int
	fmt.Print("Masukan angka 3 digit:")
	fmt.Scan(&input)
	
	digit1 = input / 100
	digit2 = (input / 100) % 10
	digit3 = input % 10
	
	fmt.Print(digit1 <= digit2 && digit2 <= digit3) // jika menggunakan operator logika seperti && atau || maka akan langsung dianggap boolean
}