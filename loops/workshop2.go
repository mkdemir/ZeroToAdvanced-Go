package loops

import "fmt"

/*
	1. Kullanıcıdan bir sayı girmesini isteyiniz.
		- 23: Asaldır.
*/
func Demo4() { // Bir paket içersinde bizim o paketteki fonksiyon isimlerimizin farklı olması gerekiyor.
	sayi := 0

	fmt.Print("Bir Sayı Giriniz: ")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 { // % mod alma
			asalMi = false
		}
	}

	if !asalMi { // asalMi == true
		fmt.Println("Asaldır.")
	} else {
		fmt.Println("Asal değildir.")
	}

}
