package main
import (
"bufio"
"fmt"
"os"
"strings"
)
func main() {
expectedColors := []string{"merah", "kuning", "hijau", "ungu"}
scanner := bufio.NewScanner(os.Stdin)
success := true
for i := 1; i <= 5; i++ {
fmt.Printf("Percobaan %d: ", i)
scanner.Scan()
colors := strings.Fields(scanner.Text())
if len(colors) != len(expectedColors) {
success = false
break
}
for j, color := range colors {
if color != expectedColors[j] {
success = false
break
}
}
if !success {
break
}
}
fmt.Printf("BERHASIL: %v\n", success)
}

package main
import (
"fmt"
"strings"
)
func main() {
var n int
var bunga, pita string
var count int
fmt.Print("Masukkan jumlah bunga (N): ")
fmt.Scan(&n)
if n == 0 {
fmt.Println("Pita: ")
fmt.Println("Bunga: 0")
return
}
for count = 1; count <= n; count++ {
fmt.Printf("Bunga %d: ", count)
fmt.Scan(&bunga)
if strings.ToUpper(bunga) == "SELESAI" {
count--
break
}
if count > 1 {
pita += " -
"
}
pita += bunga
}
fmt.Println("Pita:", pita)
fmt.Printf("Bunga: %d\n", count)
}

package main 

 

import ( 

    "fmt" 

    "math" 

) 

 

func main() { 

    var beratKantong1, beratKantong2, totalBerat float64 

 

    for { 

        fmt.Print("Masukan berat belanjaan di kedua kantong: ") 

        _, err := fmt.Scanf("%f %f", &beratKantong1, &beratKantong2) 

 

        if err != nil { 

            fmt.Println("Input tidak valid. Silakan coba lagi.") 

            continue 

        } 

 

        if beratKantong1 < 0 || beratKantong2 < 0 { 

            fmt.Println("Proses selesai.") 

            break 

        } 

 

        totalBerat = beratKantong1 + beratKantong2 

 

        if totalBerat > 150 { 

            fmt.Println("Proses selesai.") 

            break 

        } 

 

        selisih := math.Abs(beratKantong1 - beratKantong2) 

 

        if selisih >= 9 { 

            fmt.Println("Sepeda motor pak Andi akan oleng: true") 

        } else { 

            fmt.Println("Sepeda motor pak Andi akan oleng: false") 

        } 

    } 

} 

package main 

 

import ( 

    "fmt" 

    "math" 

) 

 

func f(k int) float64 { 

    numerator := math.Pow(float64(4*k+2), 2) 

    denominator := float64((4*k + 1) * (4*k + 3)) 

    return numerator / denominator 

} 

 

func sqrt2Approximation(K int) float64 { 

    product := 1.0 

    for k := 0; k <= K; k++ { 

        product *= f(k) 

    } 

    return product 

} 

 

func main() { 

    var K int 

 

    fmt.Print("Masukkan nilai K: ") 

    fmt.Scan(&K) 

    fmt.Printf("Nilai f(K) = %.10f\n", f(K)) 

 

    fmt.Print("Masukkan nilai K untuk akar 2: ") 

    fmt.Scan(&K) 

    fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2Approximation(K)) 

} 

package main 

 

import "fmt" 

 

func main() { 

    var nam float64 

    var nmk string 

    fmt.Print("Nilai akhir mata kuliah: ") 

    fmt.Scanln(&nam) 

    if nam > 80 { 

        nmk = "A" 

    } else if nam > 72.5 { 

        nmk = "AB" 

    } else if nam > 65 { 

        nmk = "B" 

    } else if nam > 57.5 { 

        nmk = "BC" 

    } else if nam > 50 { 

        nmk = "C" 

    } else if nam > 40 { 

        nmk = "D" 

    } else { 

        nmk = "E" 

    } 

    fmt.Println("Nilai mata kuliah: ", nmk) 

} 

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
