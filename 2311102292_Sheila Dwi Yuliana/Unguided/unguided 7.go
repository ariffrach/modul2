package main

import "fmt"

func main() {
	var bilangan int
	var faktor []int

	fmt.Print("Bilangan: ")
	fmt.Scanln(&bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor = append(faktor, i)
		}
	}

	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}
	fmt.Println()

	if len(faktor) == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
