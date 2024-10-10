package main

import "fmt"

func main() {
    var beratGram int

    fmt.Print("Berat parsel (gram): ")
    fmt.Scanln(&beratGram)

    beratKg := beratGram / 1000
    sisaGram := beratGram % 1000

    biayaDasar := beratKg * 10000

    var biayaTambahan int
    if sisaGram >= 500 {
        biayaTambahan = sisaGram * 5
    } else {
        biayaTambahan = sisaGram * 15
    }

    if beratKg > 10 && sisaGram < 500 {
        biayaTambahan = 0
    }

    totalBiaya := biayaDasar + biayaTambahan

    fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}