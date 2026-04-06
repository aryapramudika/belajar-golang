package main

import "fmt"

func main() {
	/*
		Konstanta adalah jenis variabel yang tidak bisa diubah setelah di deklarasikan
		Cocok untuk data yang punya nilai tetap, misal pi (22/7) dll
	*/
	const namaMertua string = "Paimin"
	const NamaBesan = "Paijo"
	fmt.Printf("Nama mertuamu adalah:\n%s\n", namaMertua)
	fmt.Printf("Nama besanmu adalah:\n%s\n", NamaBesan)

	/*
		Multi Deklarasi
	*/

	const (
		Presiden     = "Prabowo"
		Program      = "MBG"
		Rugi     int = -99999
	)
	fmt.Printf("Program bapak presiden %s, yaitu %s, rugi sebanyak %d juta dollar\n", Presiden, Program, Rugi)
}
