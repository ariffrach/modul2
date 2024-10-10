package main

import (
	"fmt"
)

func main() {
	var NAM float32
	var NMK string

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&NAM)

	if NAM >= 80 {
		NMK = "A"
	} else if NAM >= 70 {
		NMK = "B"
	} else if NAM >= 65 {
		NMK = "C"
	} else if NAM >= 45 {
		NMK = "D"
	} else if NAM >= 40 {
		NMK = "E"
	} else {
		NMK = "F"
	}

	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", NAM, NMK)
}
