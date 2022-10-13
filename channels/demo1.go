package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(cifySayiCn chan int) { // channel bir pipe benzetebilirsiniz.
	toplam := 0
	for i := 0; i <= 10; i += 2 { // i = i + 2
		toplam = toplam + i
		fmt.Println("Çift Sayı Toplama Çalışıyor")
		time.Sleep(1 * time.Second)
	}

	cifySayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 { // i = i + 2
		toplam = toplam + i
		fmt.Println("Tek Sayı Toplama Çalışıyor")
		time.Sleep(1 * time.Second)
	}

	tekSayiCn <- toplam
}
