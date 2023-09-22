package main

import (
	"fmt"
)

func hitungFrekuensi(periode float64) float64 {
	return 1.0 / periode
}

func main() {
	var periode float64

	fmt.Print("Masukkan periode getaran (sekon): ")
	fmt.Scanln(&periode)

	frekuensi := hitungFrekuensi(periode)

	fmt.Printf("Frekuensi getaran adalah: %.2f Hz\n", frekuensi)
}
