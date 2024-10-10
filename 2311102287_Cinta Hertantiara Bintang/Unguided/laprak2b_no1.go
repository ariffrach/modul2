package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)

	berhasil := true

	for c := 1; c <= 5; c++ {
		fmt.Printf("Percobaan %d: ", c)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		warna := strings.Split(input, " ")

		for i := 0; i < 4; i++ {
			if warna[i] != correctOrder[i] {
				berhasil = false
				break
			}
		}

		if !berhasil {
			break
		}
	}

	if berhasil {
		fmt.Println("Berhasil: true")
	} else {
		fmt.Println("Berhasil: false")
	}
}
