1. Source Code
package main
import "fmt"
func main() {
var nama string = "M. ARIF RACHMAN"
var umur int = 19
var tinggi float64 = 177.5
var isSunny bool = false
var inisial rune = 'A'
fmt.Println("Nama:", nama)
fmt.Println("Umur:", umur)
fmt.Println("Tinggi:", tinggi)
fmt.Println("Is sunny", isSunny)
fmt.Println("Inisial", inisial)
}

package main
import "fmt"
func main() {
//Tahun kabisat
var year int
fmt.Print("Masukkan tahun: ")
fmt.Scanf("%d", &year)
if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
fmt.Println(year, "adalah tahun kabisat (true)")
} else {
fmt.Println(year, "bukan tahun kabisat (false)")
}
}

package main
import "fmt"
func main() {
//Temperatur
var temperaturCelsius float64
fmt.Print("Temperatur Celsius: ")
fmt.Scanln(&temperaturCelsius)
var temperaturFahrenheit float64 = (temperaturCelsius * 9 / 5) + 32
var temperaturReamur float64 = temperaturCelsius * 4 / 5
var temperaturKelvin float64 = temperaturCelsius + 273.15
fmt.Printf("Derajat Fahrenheit: %.0f\n", temperaturFahrenheit)
fmt.Printf("Derajat Reamur: %.0f\n", temperaturReamur)
fmt.Printf("Derajat Kelvin: %.2f\n", temperaturKelvin)
}

package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
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
