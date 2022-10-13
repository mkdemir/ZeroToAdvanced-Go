package structs

import "fmt"

// Class'lara benzer.
// Değer güdümlü çalışır int gibi çalışır.
func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "xyz"}) // Bizim için bir obje struct yapısı
	x := product{"Bilgisayar", 500, "xyz"}
	fmt.Println(x)
	fmt.Println(product{"Bilgisayar", 500, ""})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 500})
	fmt.Println(product{name: "Bilgisayar"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
	//discountRate int // default 0'dır.
}
