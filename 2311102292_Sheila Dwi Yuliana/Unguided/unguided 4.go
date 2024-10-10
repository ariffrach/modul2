package main

import (
	"fmt"
)

func f(k int) float64 {
	return float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
}

func main() {
	var K int
	var akar2 float64

	fmt.Print("Nilai K: ")
	fmt.Scanln(&K)
	nilaiF := f(K)
	fmt.Printf("Nilai f(K) = %.10f\n", nilaiF)

	for k := 0; k <= K; k++ {
		akar2 += f(k)
	}
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
