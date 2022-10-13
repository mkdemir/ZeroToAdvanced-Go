package arrays

import "fmt"

func Demo4() {
	// Birçok programlama dilinde Jagged  array çok boyutlu dizi anlamına gelen bir matris yapısını konuştuk.
	var sayilar [2][3]int // 1.satır 2.sütun

	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[0][2] = 5
	sayilar[1][0] = 2
	sayilar[1][1] = 4
	sayilar[1][2] = 6
	fmt.Println(sayilar)
	fmt.Println(sayilar[1][1])

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
