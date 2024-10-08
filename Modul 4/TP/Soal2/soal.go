package main

import "fmt"

func main(){
	var jamkerja, upahperjam, jamlembur, gajinormal, upahlembur, total float64
	fmt.Println("Masukan jam kerja anda dalam seminggu & upah per jam:")
	fmt.Scan(&jamkerja,&upahperjam)
	if jamkerja > 40 {
		jamlembur = jamkerja - 40
		upahlembur = jamlembur * 1.5
		gajinormal = 40 * upahperjam
	} else {
		gajinormal = jamkerja * upahperjam
	}

	total = 4 * (gajinormal + upahlembur)
	fmt.Printf("Total gaji anda adalah: %.f\n ",total)
}