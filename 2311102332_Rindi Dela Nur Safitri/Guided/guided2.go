package main

import "fmt"

func main() {
    var tahun int
    fmt.Println("Program menentukan false atau true nilai tahun kabisat")
    fmt.Print("Tahun: ") 
    fmt.Scanln(&tahun)

    if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Kabisat: false")
    }

}
