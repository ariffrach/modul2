package main

import (
	"fmt"
)

func main() {
	var b int

	// Meminta input bilangan bulat dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&b)

	// Memastikan input b lebih dari 1
	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	// Mencari dan menampilkan faktor-faktor dari b
	fmt.Print("Faktor: ")
	var faktorCount int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			faktorCount++
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan tersebut merupakan bilangan prima
	if faktorCount == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
