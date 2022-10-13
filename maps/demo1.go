package maps

import "fmt"

func Demo1() {
	// key-value
	sozluk := make(map[string]string) // map yap

	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book") // Eleman silme
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["lamba"] // varMi -> False True dönüyor.
	fmt.Println(deger)
	fmt.Println("Listede Olma Durumu: ", varMi)
	// iterate etmek -> Onu gezme bir döngü ile gezmek range ile dönme

	sozluk2 := map[string]string{"glass": "bardak", "hello": "merhaba"}

	fmt.Println(sozluk2)
}
