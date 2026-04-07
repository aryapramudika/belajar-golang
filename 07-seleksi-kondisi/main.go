package main

import "fmt"

func main() {
	/*
		Berfungsi untuk mengatur alur eksekusi flow program, analoginya mirip fungsi rambu lalu lintas.

		Contoh simpel: Lampu Merah,
		Kalau Hijau boleh jalan, kalau kuning hati-hati, kalau merah berhenti
	*/
	fmt.Println("If-ElseIf-Else:")
	warnaLampu := "hijau"

	if warnaLampu == "merah" {
		fmt.Println("Sabar boss!")
	} else if warnaLampu == "kuning" {
		fmt.Println("Hati-Hati boss!")
	} else {
		fmt.Println("Silakan jalan bos!")
	}
	fmt.Println("")

	/*
		Variabel dalam kondisi, ini berlaku di kondisi saja, misal cek diskon belanja
		misal kalau harganya -> 70.000 maka dapat diskon 15%
	*/
	fmt.Println("Temporary variabel:")
	hargabarang := 60000
	if diskon := hargabarang * 15 / 100; hargabarang >= 70000 {
		fmt.Println("Anda dapat diskon: ", diskon)
		fmt.Println("Harga barang", hargabarang)
	} else {
		fmt.Println("Anda gak dapat diskon, harga barang harus lebih dari 70.000")
	}
	fmt.Println("")
	/*

		Case - Switch digunakan untuk seleksi yang sifatnya fokus pada satu variabel
		Kalau gak ada yang match maka akan balik ke kondisi default, dimana default ini di case dianggap sebagai else.
		Case-Switch pada golang, cara kerjanya berbeda dengan bahasa pemrograman pada umumnya,
		setelah menemukan yang match maka gak akan di lanjutkan ke pengecekan case selanjutnya.

		Contoh dibawah ini, cek apakah cert_k8s merupakan CKAD
	*/
	fmt.Println("Case-Switch:")

	fmt.Println("Checking apakah anda punya cert CKAD....")
	k8s_cert := "CKAD"

	switch k8s_cert {
	case "CKA":
		fmt.Println("Oke lu ada CKA")
	case "CKS":
		fmt.Println("Oke lu ada CKS")
	default:
		fmt.Println("Lu gak punya certnya")
	}
	fmt.Println("")

	/*

		Case - Switch untuk banyak kondisi
	*/

	fmt.Println("CKS_Checker...")

	cks_checker := "cks"

	switch cks_checker {
	case "cka", "ckad", "cks":
		fmt.Println("Oke anda punya CKS")
	default:
		fmt.Println("Anda gak punya CKS")
	}

	/*

		Case - Kurung kurawal di deafult bisa dipake kalau emang case defaultnya banyak statement
	*/
	fmt.Println("")
	fmt.Println("KCNA Checker...")

	kcna_checker := "kcna"

	switch kcna_checker {
	case "cka", "ckad", "cks":
		fmt.Println("Oke anda punya CKA,CKAD,CKS")
	case "kcsa":
		fmt.Println("Oke punya kcna")
	default:
		{
			fmt.Println("Anda gak punya KCNA")
			fmt.Println("Tinggal Beli Brooo")
		}
	}
}
