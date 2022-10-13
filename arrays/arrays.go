package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int //  5 elemanlı bir dizi yani içinde 5 eleman tutuyor.

	fmt.Println(sayilar) // [0 0 0 0 0] Tam sayıların default deferi 0 olarak başlatılır.

	sayilar[2] = 50 // set işlemi yaptık değeri atama

	fmt.Print(sayilar[2])
}
