package main

import "fmt"

func main() {
        var berat1, berat2 float64
        var selisih, total float64
        var oleng bool

        for {
                fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
                fmt.Scan(&berat1, &berat2)

                if berat1 < 0 || berat2 < 0 || total > 150 {
                        fmt.Println("Proses selesai.")
                        break
                }

                selisih = berat1 - berat2
                if selisih < 0 {
                        selisih = -selisih
                }
                total += berat1 + berat2

                oleng = selisih >= 9

                fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
        }
}