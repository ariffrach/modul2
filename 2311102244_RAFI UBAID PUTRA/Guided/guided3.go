package main

import "fmt"

func findFactors(b int) []int {
	var factors []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func isPrime(b int) bool {
	factors := findFactors(b)
	return len(factors) == 2
}

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	factors := findFactors(b)
	fmt.Print("Faktor: ")
	for _, factor := range factors {
		fmt.Printf("%d ", factor)
	}
	fmt.Println()

	primeStatus := isPrime(b)
	fmt.Printf("Prima: %t\n", primeStatus)
}
