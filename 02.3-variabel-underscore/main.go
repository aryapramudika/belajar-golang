package main

import "fmt"

func main() {
	/*variable underscore digunakan untuk menampung variable/nilai yang gak kepakai
	karena di golang itu by default kalau kita deklarasikan variable itu harus dipakai.

	Variable _ itu seperti blackhole ya, jadi kalau dimasukin disitu langsung hilang dan gak bisa dipanggil.
	*/
	_ = "belajar golang dengan"
	nama := "Arya Pramudika"

	fmt.Printf("Halo, %s\n", nama)
}
