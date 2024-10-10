package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal= " + satu + " " + dua + " " + tiga)
	temp = satu
	temp = dua
	temp = tiga
	tiga = temp
	fmt.Println("Output awal= " + satu + " " + dua + " " + tiga)
}
