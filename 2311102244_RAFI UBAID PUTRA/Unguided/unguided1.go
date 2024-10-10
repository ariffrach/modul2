package main

import (
	"fmt"
)

func main() {
	targetSequence := []string{"merah", "kuning", "hijau", "ungu"}
	var results []bool

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d:\n", i)
		var colors [4]string

		for j := 0; j < 4; j++ {
			fmt.Printf("%d: ", j+1)
			fmt.Scanln(&colors[j])
		}
		if checkSequence(colors[:], targetSequence) {
			results = append(results, true)
			fmt.Println("Percobaan BERHASIL : true")
		} else {
			results = append(results, false)
			fmt.Println("Percobaan BERHASIL : false")
		}
	}

	fmt.Println("\nHasil semua percobaan:")
	for i, result := range results {
		fmt.Printf("Percobaan %d: %t\n", i+1, result)
	}
}

func checkSequence(colors []string, target []string) bool {
	for i := 0; i < len(target); i++ {
		if colors[i] != target[i] {
			return false
		}
	}
	return true
}
