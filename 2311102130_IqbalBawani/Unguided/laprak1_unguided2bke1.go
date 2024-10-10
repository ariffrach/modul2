// UNGUIDED 2B KE 1
// IqbalBawani
// 2311102130
// s1if1107
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	warna := []string{"merah", "kuning", "hijau", "ungu"}
	var isCorrect = true
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan urutan warna untuk 5 percobaan (pisahkan dengan spasi):")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		if len(input) != len(warna) {
			isCorrect = false
			break
		}
		for j := 0; j < len(warna); j++ {
			if input[j] != warna[j] {
				isCorrect = false
				break
			}
		}
		if !isCorrect {
			break
		}
	}

	if isCorrect {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
