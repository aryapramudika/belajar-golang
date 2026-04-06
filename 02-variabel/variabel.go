package main

import "fmt"

func main() {
	//inisialiasi variabel dengan tipe data, dan isi datanya langsung
	var namaDepan string = "Arya"

	//inisialiasi variable dengan tipe data string, datanya belakangan
	var namaBelakang string

	//ini datanya
	namaBelakang = "Pramudika"

	/*
		pakai printf dimana dengan ini perlu mendefinisikan tipe datanya %s, artinya string
		jadi karena gw mau print namaDepan dan namaBelakang maka %s %s
		lalu gw tambahin \n diakhir baris untuk menambah baris baru karena by default printf ini tidak membuat
		baris baru seperti Println
	*/
	fmt.Printf("Halo mas %s %s !\n", namaDepan, namaBelakang)
}
