package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo1()
	//conditionsals.Demo2()
	//conditionals.Demo3()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//loops.Demo4()
	//loops.Demo5()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo4()
	//slices.Demo2()

	/*
		functions.SelamVer("Mustafa")
		functions.SelamVer("Kaan")
		// functions.Topla(2, 3) //parametre diyoruz.

		var sonuc = functions.Topla(2, 6)
		fmt.Println(sonuc * 10)
	*/

	/*
		sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6) // _ -> Bana sonucu verme diyorsunuz.

		fmt.Println("Toplam: ", sonuc1)
		fmt.Println("Çıkarma: ", sonuc2)
		fmt.Println("Çarpma: ", sonuc3)
	*/
	//fmt.Println("Bölme: ", sonuc4)

	// var sonuc = functions.ToplaVariadic(1, 4, 5, 6, 3, 10)
	// fmt.Println()
	// ctrl + k ctrl + c

	// fmt.Println(functions.ToplaVariadic(1, 4, 5, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic()) // 0

	// sayilar := []int{4, 6, 7, 8, 11}

	// fmt.Println(functions.ToplaVariadic(sayilar...))
	// 	maps.Demo1()
	//for_range.Demo1()

	//for_range.Demo2()
	//for_range.Demo3()
	// sayi := 20            // Değer tiplerdir. (Değeri tutan tiplerdir) (Sayısal değerler) - int, float, bool
	// pointers.Demo1(&sayi) // pointer çağırıken (Yollarken &) * Adres anlamına geliyor.
	// fmt.Println("Maindeki sayı :", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayı: ", sayilar[0]) // Array'ler değerleri üzerinden çalışmaz array'ler referansları göndeririz. - map, sliece gibi yapılar.

	// structs.Demo2(

	// go goroutines.CiftSayilar() // threat başka
	// go goroutines.TekSayilar()

	// time.Sleep(5 * time.Second) // time.Second -> 1 saniyeye karşılık gelir.(Yukarıdakiler tam çalışacaktı ama zaman kalmadı) Treat'lere yetişemedi

	// fmt.Println("Main Bitti")

	// go rutine: channel
	// ciftSayiCn := make(chan int) // channel'ları tanımladık
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn) // threat başka (Parametre olarak yolladık)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn // Okuma

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	// interfaces

	// interfaces.Demo1()
	//interfaces.Demo2()

	// defer_statement.B()
	//defer_statement.Demo3()

	// error_handling.Demo1()
	// interfaces.Demo3()

	//fmt.Println(error_handling.TahminEt2(102))
	// string_functions.Demo1()
	//string_functions.Demo2()

	//restful.Demo2()

	//project.GetAllProducts()
	// project.AddProduct()
	// project.GetAllProducts()
	product, _ := project.AddProduct()
	fmt.Println(product)
	products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
