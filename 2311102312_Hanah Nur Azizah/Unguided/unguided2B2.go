package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//HANAH NUR AZIZAH (2311102312)
func main() {
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("N: ")
	var N int
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		N, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil || N <= 0 {
			fmt.Println("Masukkan bilangan bulat positif")
		} else {
			break
		}
	}

	var pita string
	var count int

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		input = strings.TrimSpace(input)

		if strings.ToUpper(input) == "SELESAI" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita = pita + " – " + input
		}
		count++
	}

	fmt.Println("Pita:", pita)
	fmt.Printf("Bunga: %d\n", count)
}
