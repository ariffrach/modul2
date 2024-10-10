package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu) // Masukan kata pertama
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)  // Masukan kata kedua
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga) // Masukan kata ketiga
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	
	// Swapping the strings
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}