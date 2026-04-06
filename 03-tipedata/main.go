package main

import "fmt"

func main() {
	//tipe data numerik/non desimal
	/*
		uint: tipe data bilangan cacah (bilangan positif)
		int: tipe data untuk bilangan bulat (negatif dan positif)

		Kedua bilangan tersebut terbagi menjadi beberapa jenis:

		uint8: 0-255
		uint16: 0-65535
		uint32: 0-4294967295
		uint64: 0-18446744073709551615
		uint: sama dengan uint32 atau uint64 (tergantung nilai)
		byte: sama dengan uint8
		int8: -128-127
		int16: -32768-32767
		int32: -2147483648-2147483647
		int64: -9223372036854775808-9223372036854775807
		int: sama dengan int32 atau int64 (tergantun nilai)
		rune: sama dengan int32

		Untuk menentukan tipe data variabel gak boleh sembarangan bro, karena nanti berkaitan dengan alokasi memori variabel.
		Jadi kalau lu nentuinnya bener nanti programnya optimal, gak makai memori berlebih.
	*/
	var GatewayAddress uint8 = 254
	var HutangBayarWoy = -1000000

	fmt.Println("Ini Tipe data Numerik dan Non Desimal")
	fmt.Printf("IP Gateway Router yaitu 192.168.1.%d\n", GatewayAddress)
	fmt.Printf("Hutang yang kamu perlu bayar yaitu: %d\n", HutangBayarWoy)
	fmt.Println("")

	//tipe data desimal nih
	/*
		ada float32 dan float64 cakupannya lebar bisa baca di:
		http://www.h-schmidt.net/FloatConverter/IEEE754.html
	*/
	var contohNumberBro = 79.9

	fmt.Println("Ini tipe data desimal")
	//format %f untuk mengubah format desimal menjadi string, digit desimalnya yaitu 6 digit. Contoh 79.9 jadi 79.900000
	fmt.Printf("Sisa Duitmu Rp.%f\n", contohNumberBro)
	//kalau mau kontrol jumlah digitnya tinggal diubah dengan format %.nf, contoh ini %.1f maka jadi 79.9
	fmt.Printf("Nilai Ujianmu adalah: %.1f\n", contohNumberBro)
	fmt.Println("")

	//tipe data boolean (bool)
	/*
		variable boolean isinya cuma 2 yaitu kalau nggak true atau false
		Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan.
	*/
	fmt.Println("Boolean")
	var punyaPacar bool = true
	//%t untuk format boolean
	fmt.Printf("Apakah anda punya pacar? %t \n", punyaPacar)
	fmt.Println("")

	//tipe data string
	/*
		Ciri khasnya yaitu diapit "", contoh penerapan langsung aja bro
	*/
	fmt.Println("String")

	//string biasa
	var suratCinta string = "Halo, aku cinta padamu!"

	//backtick bisa digunain untuk bikin string multiline juga untuk escape misal mau masukin "" di stringnya
	var untukSiapa string = `Tentunya untuk cewe cantik, siapa emangnya?
	Yaitu "xxx"
	`
	fmt.Printf("Ada surat nih:\n%s\n", suratCinta)
	fmt.Println(untukSiapa)

}
