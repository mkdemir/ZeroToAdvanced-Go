package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func D() {
	fmt.Println("D fonksiyonu çalıştı")
}

func B() {
	defer A() // Last in first out (java'da C# Finally ile eş tutabiliriz).
	defer C()
	defer D()                           // İkinci by çalışıyor.
	fmt.Println("B fonksiyonu çalıştı") // İlk bu çalışıyor.
	// defer ile uygulmanın sonunda onun çalışmasını garanti ediyoruz.
}
