package main

import "fmt"

func main() {
	var (
		a          int
		c, f, r, k float64
	)

	fmt.Print("Pilih lah program untuk mengkonversi suhu \n")
	fmt.Print("1. Dari Celcius\n2. Dari Fahrenheit\n3. Reamur\n4.Kelvin\n\n")
	fmt.Print("Masukan 1 / 2 / 3 / 4 : ")
	fmt.Scanf("%d", &a)
	if a == 1 {
		fmt.Print("Masukan Nilai Celcius : ")
		fmt.Scanf("%f", &c)
		fahrenheitc := (c * 1.8) + 32
		reamur := c * 0.8
		kelvinr := c + 273.15
		fmt.Print("Nilai Konversi dari Celcius ke Fahrenheit : ", fahrenheitc)
		fmt.Print("\nNilai Konversi dari Celcius ke Reamur : ", reamur)
		fmt.Print("\nNilai Konversi dari Celcius ke Kelvin : ", kelvinr)
	} else if a == 2 {
		fmt.Print("Masukan Nilai Fahrenheit : ")
		fmt.Scanf("%f", &f)
		celciusf := (f - 32) / 1.8
		reamurf := (f - 32) / 2.25
		kelvinf := (f + 459.67) / 1.8
		fmt.Print("Nilai Konversi dari Fahrenheit ke Celcius : ", celciusf)
		fmt.Print("\nNilai Konversi dari Fahrenheit ke Reamur : ", reamurf)
		fmt.Print("\nNilai Konversi dari Fahrenheit ke Kelvin : ", kelvinf)
	} else if a == 3 {
		fmt.Print("Masukan Nilai Reamur : ")
		fmt.Scanf("%f", &r)
		celciusr := r / 0.8
		fahreinheitr := (r * 2.25) + 32
		kelvinr := (r / 0.8) + 273.15
		fmt.Print("Nilai Konversi dari Reamur ke Celcius : ", celciusr)
		fmt.Print("\nNilai Konversi dari Reamur ke Fahrenheit : ", fahreinheitr)
		fmt.Print("\nNilai Konversi dari Reamur ke Kelvin : ", kelvinr)
	} else if a == 4 {
		fmt.Print("Masukan Nilai Kelvin : ")
		fmt.Scanf("%f", &k)
		celciusk := k - 273.15
		fahreinheitk := (k * 1.8) - 459.67
		reamurk := (k - 273.15) * 0.8
		fmt.Print("Nilai Konversi dari Kelvin ke Celcius : ", celciusk)
		fmt.Print("\nNilai Konversi dari Kelvin ke Fahrenheit : ", fahreinheitk)
		fmt.Print("\nNilai Konversi dari Kelvin ke Reamur : ", reamurk)
	} else {
		fmt.Print("\nSilahkan Masukan Input Sesuai Menu yang Disediakan")
	}
}
