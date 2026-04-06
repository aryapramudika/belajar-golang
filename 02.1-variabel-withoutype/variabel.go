package main

import "fmt"

func main() {
	var namaLengkap string = "Arya Pramudika"

	/*ini contoh variabel tapi gak definisiin si tipe datanya, golang bisa tau dari isinya
	menggunakan tanda := untuk cara ini, dan tidak perlu mendefinisikan var dan =
	*/
	umur := 20

	fmt.Printf("Nama: %s \nUmur: %d \n", namaLengkap, umur)
}
