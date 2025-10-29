package main

import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&a)
	Perkalian := a * 2
	fmt.Printf("Hasil perkalian angka yang anda masukkan dengan 2 adalah %d\n", Perkalian)
}
