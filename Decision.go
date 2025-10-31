package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai Anda (0-100): ")
	fmt.Scan(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Selamat, anda telah lulus kkm")
	} else if nilai > 100 || nilai < 0 {
		fmt.Println("Nilai yang anda masukkan tidak valid")
	} else {
		fmt.Println("Maaf, anda belum lulus kkm")
	}

}
