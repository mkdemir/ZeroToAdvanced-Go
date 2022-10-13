package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap { // go'nun derlenen halinde herhangi bir () olmuyor.
		fmt.Print("Hesaptaki para yetersiz")
	}
	if cekilmekIstenen <= hesap {
		fmt.Print("Paranız hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("Bitti. Hesaptaki Para: " + fmt.Sprintf("%v", hesap)) // Metinsel bir ifadeyi sayı ile toplmaya çalışıyorsun senin tip dönüşümü yapman lazım.
	// .Sprintf("%f", hesap) -> Formatla float formatında ondalıklı göreceksiniz. Biz genellikle %v kullanırız. Value'dan geliyor.

	fmt.Printf("Bitti. Hesaptaki Para: %v", hesap)

	var durum bool

	durum = cekilmekIstenen > hesap

	if durum {
		fmt.Print("Hesaptaki para yetersiz")
	}
	if !durum {
		fmt.Print("Paranız hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("Bitti. Hesaptaki Para: %v", durum)
}
