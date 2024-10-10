package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	sqrt2 := 1.0
	for k := 0; k <= K; k++ {
		k := float64(k)
		hasil := (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))
		sqrt2 *= hasil
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}
