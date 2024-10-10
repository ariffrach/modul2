// UNGUIDED 2C KE 3
// IQBAL BAWANI
// 2311102130
// S1IF1107

package main

import (
	"fmt"
)

func main() {
	var Iqbalbawani int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&Iqbalbawani)
	fmt.Print("Faktor: ")
	for i := 1; i <= Iqbalbawani; i++ {
		if Iqbalbawani%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	prima := true
	if Iqbalbawani <= 1 {
		prima = false
	} else {
		for i := 2; i*i <= Iqbalbawani; i++ {
			if Iqbalbawani%i == 0 {
				prima = false
				break
			}
		}
	}
	fmt.Println("Prima:", prima)
}
