package main

import "fmt"

func main() {
	var a int

	//contoh deklarasi multi variabel dengan tipe data string
	var b, c string

	//contoh lebih ringkas, tanpa declare var dan tipe data
	d, e := "Madiun", "Samarinda"

	//contoh lebih ringkas dengan tipe data beda-beda
	f, g := 99.5, 100

	//ini isi datanya
	b, c = "Desember", "2003"
	a = 13

	/*
		Ini gw print variable a dengan tipe data integer, dan variabel b, c dengan tipe data string
	*/
	fmt.Printf("Tanggal lahir:\n%d %s %s\nTTL: %s %s\n", a, b, c, d, e)
	fmt.Printf("Nilai Ujian Matematika:\n%f\nNilai Ujian Bahasa Inggris:\n%d\n", f, g)
}
