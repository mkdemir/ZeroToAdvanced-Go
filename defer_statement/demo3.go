package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kayıtedildi: ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	// defer p.save() Ürün kayıtedildi:  Laptop
	p = product{productName: "Mouse", unitPrice: 45}
	defer p.save() // Ürün kayıtedildi:  Mouse
	fmt.Println("İşlem Başarılı")
	fmt.Println(p.productName)
}
