package main

import (
	"fmt"
	"math"
)

func main() {
	var kantongTerpal1, kantongTerpal2 float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantongTerpal1, &kantongTerpal2)
		total := kantongTerpal1 + kantongTerpal2

		if total > 150 {
			fmt.Println("Total berat tidak boleh melebihi 150 kg")
			fmt.Println("Program Selesai")
			break

		} else if kantongTerpal1 < 0 || kantongTerpal2 < 0 {
			fmt.Println("Berat tidak boleh negatif")
		} else if math.Abs(kantongTerpal1-kantongTerpal2) < 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		}
	}
}
