package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"
	//fmt.Println(metin) // ln (line) -> Satır demek.

	// go dilinde sadece for döngüsünü var. While ile aynı yapıda göndüler kullanılabiliyor.

	i := 1 // index = sayaç demekdir. Başlangıç değerini ve o döngünün sonraki eleman sayısını tutmaya yarar.
	// infinite loop : Sonsuz döngüye girer.
	// snippets
	/*
		for i <= 5 { // blok {} if gibi şu şart sağlandığı sürece bu bloğun içini çalıştır.

			fmt.Println(metin)
		}

	*/
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1 // 1 2 3 4 5 6(Durdu))
	}

	fmt.Println("Bitti")
}
