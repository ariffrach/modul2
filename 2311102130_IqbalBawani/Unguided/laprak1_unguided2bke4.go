// UNGUIDED 2B KE 4
// IQBAL BAWANI
// 2311102130
// S1IF1107
package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Print("Masukan Nilai K = ")
	fmt.Scan(&K)

	Iqbalbawani := 1.0
	for k := 0; k <= K; k++ {
		k := float64(k)
		hasil := (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))
		Iqbalbawani *= hasil
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", Iqbalbawani)
}
