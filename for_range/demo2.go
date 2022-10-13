package for_range

import "fmt"

// Workshop 1-10 arasındaki sayılardan tek olanları toplayan program
func Demo2() {
	/*
		sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		total := 0
		for _, sayi := range sayilar {
			if sayi%2 != 0 {
				total = total + sayi
			}
		}
	*/
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0

	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			toplam = toplam + sayi
		}
	}

	fmt.Println("Toplam:", toplam)

}
