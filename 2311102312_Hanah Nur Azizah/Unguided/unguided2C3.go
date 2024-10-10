package main
import (
	"fmt"
)
//HANAH NUR AZIZAH 2311102312
func main() {
	var b int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)

	if b <= 0 {
		fmt.Println("Bilangan harus lebih besar dari 0.")
		return
	}

	fmt.Printf("Bilangan: %d\nFaktor: ", b)
	factors := []int{} 

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			factors = append(factors, i)
			if i == b {
				fmt.Print(i)
			} else {
				fmt.Print(i, ", ")
			}
		}
	}

	isPrime := len(factors) == 2

	fmt.Printf("\nPrima: %t\n", isPrime)
}
