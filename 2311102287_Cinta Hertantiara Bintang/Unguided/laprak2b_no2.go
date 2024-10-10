package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bunga := []string{}
	reader := bufio.NewReader(os.Stdin)

	for perulangan := 1; ; perulangan++ {
		fmt.Printf("Bunga %d : ", perulangan)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "SELESAI" {
			break
		}

		bunga = append(bunga, input)
	}

	result := strings.Join(bunga, " - ")
	fmt.Println("Pita : " + result)
	fmt.Printf("Bunga : %d\n", len(bunga))
}
