// package main

// import (
// 	"fmt"
// )

// type DataEksperimen [5][4]string

// func InputDataEksperimen(data *DataEksperimen) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("Percobaan %d:\n", i+1)
// 		for j := 0; j < 4; j++ {
// 			fmt.Scan(&data[i][j])
// 		}
// 	}
// }

// func IsValid(data DataEksperimen, warnaYangDiHarapkan [4]string) bool {
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 4; j++ {
// 			if data[i][j] != warnaYangDiHarapkan[j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func PrintHasil(isValid bool) {
// 	if isValid {
// 		fmt.Println("BERHASIL: true")
// 	} else {
// 		fmt.Println("BERHASIL: false")
// 	}
// }

// func main() {
// 	var dataEksperimen DataEksperimen
// 	warnaYangDiHarapkan := [4]string{"merah", "kuning", "hijau", "ungu"}

// 	InputDataEksperimen(&dataEksperimen)
// 	isValid := IsValid(dataEksperimen, warnaYangDiHarapkan)
// 	PrintHasil(isValid)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"os"
// )

// func main() {
// 	var n int
// 	var bunga, pita string
// 	var count int

// 	fmt.Fprint(os.Stdout, "Masukkan jumlah bunga (N): ")
// 	fmt.Fscan(os.Stdin, &n)

// 	if n == 0 {
// 		fmt.Fprintln(os.Stdout, "Pita: ")
// 		fmt.Fprintln(os.Stdout, "Bunga: 0")
// 		return
// 	}

// 	for count < n {
// 		count++
// 		fmt.Fprintf(os.Stdout, "Bunga %d: ", count)
// 		fmt.Fscan(os.Stdin, &bunga)

// 		if strings.EqualFold(bunga, "SELESAI") {
// 			count--
// 			break
// 		}

// 		if len(pita) > 0 {
// 			pita += " - "
// 		}
// 		pita += bunga
// 	}

// 	fmt.Fprintf(os.Stdout, "Pita: %s\n", pita)
// 	fmt.Fprintf(os.Stdout, "Bunga: %d\n", count)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var beratKantong1, beratKantong2, totalBerat float64

// 	for {
// 		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
// 		_, err := fmt.Scanf("%f %f", &beratKantong1, &beratKantong2)
// 		if err != nil {
// 			fmt.Println("Input tidak valid. Silakan coba lagi.")
// 			continue
// 		}
// 		if beratKantong1 < 0 || beratKantong2 < 0 {
// 			fmt.Println("Proses selesai.")
// 			break
// 		}
// 		totalBerat = beratKantong1 + beratKantong2

// 		if totalBerat > 150 {
// 			fmt.Println("Proses selesai.")
// 			break
// 		}

// 		selisih := math.Abs(beratKantong1 - beratKantong2)

// 		if selisih >= 9 {
// 			fmt.Println("Sepeda motor pak Andi akan oleng: true")
// 		} else {
// 			fmt.Println("Sepeda motor pak Andi akan oleng: false")
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func f(k int) float64 {
// 	numerator := math.Pow(float64(4*k+2), 2)
// 	denominator := float64((4*k + 1) * (4*k + 3))
// 	return numerator / denominator
// }

// func sqrt2Approximation(K int) float64 {
// 	product := 1.0
// 	for k := 0; k <= K; k++ {
// 		product *= f(k)
// 	}
// 	return product
// }

// func main() {
// 	var K int

// 	fmt.Print("Masukkan nilai K: ")
// 	fmt.Scan(&K)
// 	fmt.Printf("Nilai f(K) = %.10f\n", f(K))

// 	fmt.Print("Masukkan nilai K untuk akar 2: ")
// 	fmt.Scan(&K)
// 	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2Approximation(K))
// }

// package main

// import "fmt"

// func main() {
// 	var berat, kg, gr, biayaKirim, tambahanBiaya, totalBiaya int

// 	fmt.Print("Berat parsel (gram): ")
// 	fmt.Scan(&berat)

// 	kg = berat / 1000
// 	gr = berat % 1000

// 	if kg > 10 {
// 		tambahanBiaya = 0
// 	} else if gr >= 500 {
// 		tambahanBiaya = gr * 5
// 	} else {
// 		tambahanBiaya = gr * 15
// 	}

// 	biayaKirim = kg * 10000
// 	totalBiaya = biayaKirim + tambahanBiaya

// 	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
// 	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKirim, tambahanBiaya)
// 	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
// }

// package main

// import "fmt"

// func main() {
// 	var nam float64
// 	var nmk string
// 	fmt.Print("Nilai akhir mata kuliah: ")
// 	fmt.Scanln(&nam)
// 	if nam > 80 {
// 		nmk = "A"
// 	} else if nam > 72.5 {
// 		nmk = "AB"
// 	} else if nam > 65 {
// 		nmk = "B"
// 	} else if nam > 57.5 {
// 		nmk = "BC"
// 	} else if nam > 50 {
// 		nmk = "C"
// 	} else if nam > 40 {
// 		nmk = "D"
// 	} else {
// 		nmk = "E"
// 	}
// 	fmt.Println("Nilai mata kuliah: ", nmk)
// }

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
