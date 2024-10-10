// UNGUIDED 2C KE 2
// IQBAL BAWANI
// 2311102130
// S1IF1107

package main

import (
	"fmt"
)

func main() {
	var nam float64
	var nmk string

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam)

	if nam >= 80 {
		nmk = "A"
	} else if nam >= 70 {
		nmk = "B"
	} else if nam >= 60 {
		nmk = "C"
	} else if nam >= 45 {
		nmk = "D"
	} else if nam >= 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}

	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
