//UNGUIDED 2B KE 2
//IQBALBAWANI
//2311102130
//S1IF1107

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga yang akan dimasukkan: ")
	fmt.Scan(&N)

	bunga := []string{}
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToUpper(input) == "SELESAI" {
			break
		}

		if input != "" {
			bunga = append(bunga, input)
		} else {
			fmt.Println("Nama bunga tidak boleh kosong!")
			i--
		}
	}

	result := strings.Join(bunga, " – ")
	fmt.Println("Pita : " + result)
	fmt.Printf("Bunga : %d\n", len(bunga))
}
