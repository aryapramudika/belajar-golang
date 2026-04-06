package main

import "fmt"

func main() {
	/*
		Operator Aritmatika
		-------------------
		+ (penjumlahan)
		- (pengurangan)
		* (perkalian)
		/ (pembagian)
		% (modulus/sisa hasil pembagian)
		-------------------
	*/
	fmt.Println("Operator Aritmatika:")
	nilai := 90
	fmt.Println("Nilai ujianmu adalah:", nilai)
	nilai2 := nilai - 20
	fmt.Println("Nilai ujian adit", nilai2)
	nilai3 := nilai / 2
	fmt.Println("Nilai ujian agus", nilai3)
	nilai4 := nilai % 4
	fmt.Println("Nilai ujian agung", nilai4)
	fmt.Println("")

	fmt.Println("Operator Perbandingan:")
	/*
		Operator Perbandingan
		---------------------
		== (apakah nilai kiri sama dengan nilai kanan)
		!= (apakah nilai kiri tidak sama dengan nilai kanan)
		<  (apakah nilai kiri lebih kecil dari nilai kanan)
		<= (apakah nilai kiri lebih kecil atau sama dengan nilai kanan)
		>  (apakah nilai kiri lebih besar dari nilai kanan)
		>= (apakah nilai kiri lebih besar sama dengan nilai kanan)
		---------------------
	*/

	/*
		Disini casenya ada variable contohAngka, dimana nanti dibandingkan apakah sama dengan variabel perbandinganAngka
		kalau sama hasilnya true, kalau tidak sama hasilnya false
	*/
	contohAngka := (2 + 6)
	perbandinganAngka := contohAngka == 8
	fmt.Printf("Nilai dari contohAngka merupakan: %d (%t)\n", contohAngka, perbandinganAngka)

	/*
		Disini casenya ada variable contohAngka2, dimana nanti jika dibagi 2 apakah sisa 0 (bilangan genap)
		kalau sisa 0 berarti hasilnya true, kalau sisa > 0 maka hasilnya false
	*/
	contohAngka2 := (100)
	perbandinganAngka2 := contohAngka2%2 == 0
	fmt.Printf("Apakah nilai dari vaiabel contohAngka2 bilangan genap?, nilai contohAngka2 yaitu: %d, itu termasuk bilangan genap=(%t)\n", contohAngka2, perbandinganAngka2)
	fmt.Println("")

	fmt.Println("Operator Logika:")
	/*
		Operator Perbandingan
		---------------------
		&& (kiri dan kanan) AND
		|| (kiri atau kanan) OR
		!  (negasi/nilai kebalikan) NOT
		---------------------
		Casenya digunakan untuk membandingkan tipedata boolean
	*/

	contohNilai1 := true
	contohNilai2 := false

	perbandinganNilai1 := contohNilai1 && contohNilai2
	fmt.Printf("Isi dari contohNilai1(%t) dan ContohNilai2(%t) jika dibandingkan(AND &&) hasilnya: \t (%t)\n", contohNilai1, contohNilai2, perbandinganNilai1)

	perbandinganNilai2 := contohNilai1 || contohNilai2
	fmt.Printf("Isi dari contohNilai1(%t) dan contohNilai2(%t) jika dibandingkan (OR ||) hasilnya: \t (%t)\n", contohNilai1, contohNilai2, perbandinganNilai2)

	perbandinganNilai3 := !contohNilai1
	fmt.Printf("Isi dari contohNilai1(%t), jika dibalik (!), maka hasilnya: \t (%t)\n", contohNilai1, perbandinganNilai3)

}
