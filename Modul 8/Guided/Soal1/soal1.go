package main

import "fmt"

func main() {
    var bilangan int

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    if bilangan < 0 {
        bilangan = -bilangan
    }

    fmt.Printf("Nilai mutlak dari bilangan tersebut adalah: %d\n", bilangan)
}