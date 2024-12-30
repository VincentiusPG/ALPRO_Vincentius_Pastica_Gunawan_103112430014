package main

import "fmt"

func main() {
    var kapasitas int
    var isiSekarang int = 0
    var volume int
    
    // Input kapasitas tangki
    fmt.Print("Masukkan kapasitas tangki (T): ")
    fmt.Scan(&kapasitas)
    
    // Perulangan untuk mengisi tangki
    for isiSekarang < kapasitas {
        // Input volume ember
        fmt.Print("Masukkan volume air dalam ember (V): ")
        fmt.Scan(&volume)
        
        // Tambahkan volume ke isi sekarang
        isiSekarang += volume
        
        // Cek apakah tangki sudah penuh
        if isiSekarang >= kapasitas {
            fmt.Println("true")  // Tangki penuh
        } else {
            fmt.Println("false") // Tangki belum penuh
        }
        
        // Tampilkan status pengisian
        fmt.Printf("Status: %d/%d liter\n", isiSekarang, kapasitas)
        fmt.Println("------------------------")
    }
}

//Pseudocode
// program TangkiAir
//     kamus
//         kapasitas, isiSekarang, volume: integer
//     algoritma
//         print "Masukkan kapasitas tangki (T): "
//         read kapasitas

//         isiSekarang = 0

//         repeat
//             print "Masukkan volume air dalam ember (V): "
//             read volume

//             isiSekarang = isiSekarang + volume

//             if isiSekarang >= kapasitas then
//                 print "true" 
//             else
//                 print "false" 
//             end if

//             print "Status: ", isiSekarang, "/", kapasitas, " liter"
//             print "------------------------"
//         until isiSekarang >= kapasitas
//     end program