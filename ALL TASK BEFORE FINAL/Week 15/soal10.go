package main

import "fmt"

func main() {
	var y int
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	var a1, a2, a3, a4, a5, a6, a7, a8, a9 int
	fmt.Println("Masukkan 9 bilangan satu per satu:")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9)

	total := a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9

	minimumSum := y * 5

	var median int

	if total >= minimumSum {
		median = y
	} else {
		median = 0
	}

	fmt.Println("Median bernilai :", median)
}

//Pseudocode
// program Median

//     kamus

//         y: integer
//         a1, a2, a3, a4, a5, a6, a7, a8, a9: integer
//         total: integer
//         median: integer
//         minimumSum: integer

//     algoritma
    
//         input(y)  // Input the value of y
//         input(a1, a2, a3, a4, a5, a6, a7, a8, a9)  
//         total = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 
//         minimumSum = y * 5  

//         if total >= minimumSum then  
//             median = y
//         else
//             median = 0
//         endif
//         output(median)  

// endprogram