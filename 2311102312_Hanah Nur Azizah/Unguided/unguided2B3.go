package main
import (
	"fmt"
)
//HANAH NUR AZIZAH 2311102312
func main() {
	var beratKantongKiri, beratKantongKanan float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong (misal: 50 60): ")
		fmt.Scan(&beratKantongKiri, &beratKantongKanan)

		if beratKantongKiri < 0 || beratKantongKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat := beratKantongKiri + beratKantongKanan
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisihBerat := beratKantongKiri - beratKantongKanan
		if selisihBerat < 0 {
			selisihBerat = -selisihBerat
		}

		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
