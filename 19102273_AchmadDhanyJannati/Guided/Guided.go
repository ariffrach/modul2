package main

import "fmt"

func main() {
    // Variabel awal
    var nama string = "Achmad Dhany Jannati"
    var umur int = 23
    var tinggi float64 = 175
    var isSunny bool = false
    var inisial rune = 'D'

    // Menampilkan variabel
    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Tinggi:", tinggi)
    fmt.Println("Is Sunny:", isSunny)
    fmt.Println("Inisial:", inisial)

    // Input dari pengguna
    var satu, dua, tiga, temp string
    fmt.Print("Masukkan input string (1): ")
    fmt.Scanln(&satu)
    fmt.Print("Masukkan input string (2): ")
    fmt.Scanln(&dua)
    fmt.Print("Masukkan input string (3): ")
    fmt.Scanln(&tiga)

    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

    // Menukar nilai
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp

    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)

    // Program untuk mengecek tahun kabisat
    var tahun int
    fmt.Println("Program menentukan apakah tahun kabisat (true/false)")
    fmt.Print("Masukkan tahun: ")
    fmt.Scanln(&tahun)

    if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

    // Perulangan 5 kali
    for i := 0; i < 5; i++ {
        fmt.Println("hello world")
    }
}
