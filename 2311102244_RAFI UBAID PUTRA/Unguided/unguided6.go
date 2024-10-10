package main

import (
	"fmt"
	"math"
)

func f(k int) float64 {
	numerator := math.Pow(float64(4*k+2), 2)
	denominator := float64((4*k + 1) * (4*k + 3))
	return numerator / denominator
}

func sqrt2Approximation(K int) float64 {
	product := 1.0
	for k := 0; k <= K; k++ {
		product *= f(k)
	}
	return product
}

func main() {
	var K int

	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)
	fmt.Printf("Nilai f(K) = %.10f\n", f(K))

	fmt.Print("Masukkan nilai K untuk akar 2: ")
	fmt.Scan(&K)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2Approximation(K))
}
