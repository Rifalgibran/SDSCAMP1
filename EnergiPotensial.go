package main

import (
	"fmt"
)

const gravitasi = 9.81

func hitungEnergiPotensial(massa, tinggi float64) float64 {
	return massa * gravitasi * tinggi
}

func main() {
	var massa, tinggi float64

	fmt.Print("Masukkan massa objek (kg): ")
	fmt.Scanln(&massa)

	fmt.Print("Masukkan tinggi objek di atas permukaan tanah (m): ")
	fmt.Scanln(&tinggi)

	energiPotensial := hitungEnergiPotensial(massa, tinggi)

	fmt.Printf("Energi potensial objek adalah: %.2f Joule\n", energiPotensial)
}
