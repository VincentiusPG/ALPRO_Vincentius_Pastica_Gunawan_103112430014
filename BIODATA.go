package main

import "fmt"

func main() {
 var (
 Nama, NIM, Kelas string

 )
 fmt.Print("Masukan input Nama: ")
 fmt.Scanln(&Nama)
 fmt.Print("Masukan input NIM: ")
 fmt.Scanln(&NIM)
 fmt.Print("Masukan input Kelas: ")
 fmt.Scanln(&Kelas)
 fmt.Println("Perkenalkan nama saya adalah " + Nama + ", salah satu mahasiswa ProdiS1-IF dari kelas " + Kelas + "dengan NIM " + NIM + ".")
}