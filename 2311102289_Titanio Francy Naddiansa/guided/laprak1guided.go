// package main

// import "fmt"

// func main() {
// 	var nama string = "Zein"
// 	var umur int = 20
// 	var tinggi float64 = 175.5
// 	var isSunny bool = false
// 	var inisial rune = 'A'
// 	fmt.Println("Nama:", nama)
// 	fmt.Println("Umur:", umur)
// 	fmt.Println("Tinggi:", tinggi)
// 	fmt.Println("is sunny", isSunny)
// 	fmt.Println("inisial", inisial)
//}

// package main

// import (
// 	"fmt"
// )

// func cekTahunKabisat(tahun int) bool {
// 	if tahun%400 == 0 {
// 		return true
// 	} else if tahun%100 == 0 {
// 		return false
// 	} else if tahun%4 == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	var tahun int
// 	fmt.Print("Tahun: ")
// 	fmt.Scan(&tahun)

// 	kabisat := cekTahunKabisat(tahun)

// 	fmt.Printf("Tahun: %d Kabisat: %t\n", tahun, kabisat)
// }

// package main

// import "fmt"

// func main() {
// 	var temperaturCelsius float64
// 	fmt.Print("Temperatur Celsius: ")
// 	fmt.Scanln(&temperaturCelsius)

// 	var temperaturFahrenheit float64 = (temperaturCelsius * 9 / 5) + 32
// 	var temperaturReamur float64 = temperaturCelsius * 4 / 5
// 	var temperaturKelvin float64 = temperaturCelsius + 273.15

// 	fmt.Printf("Derajat Fahrenheit: %.0f\n", temperaturFahrenheit)
// 	fmt.Printf("Derajat Reamur: %.0f\n", temperaturReamur)
// 	fmt.Printf("Derajat Kelvin: %.2f\n", temperaturKelvin)
// }

package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
