package main

import "fmt"

func main(){
	var input1 ,input2 int

	fmt.Scan(&input1,&input2)

	for done := false ; !done ; {
		input1 = input1 - input2
		fmt.Println(input1)
		done = input1 <= 0
	}
	fmt.Println(input1 == 0)
}