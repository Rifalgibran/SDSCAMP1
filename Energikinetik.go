package main

import "fmt"

func main() {
	var m, v, k float64

	fmt.Printf("\n Masukkan nilai masa benda (kg) = ")
	fmt.Scan(&m)
	fmt.Printf("\n Masukkan kecepatan benda (m/s) =")
	fmt.Scan(&v)
	k = ((v * v) * m) / 2
	fmt.Printf("Energi Kinetik = %.f", k)

}
