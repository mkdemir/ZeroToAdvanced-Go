package interfaces

import "fmt"

type Mortagage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortagage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 { // Aylık Kredi Hesaplama
	// Birden fazla kredi olabilir.
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate() // polymorphism
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortagage{rate: 10, creditPaymentTotal: 100000, address: "İstanbul"}
	credit2 := Mortagage{rate: 12, creditPaymentTotal: 50000, address: "Ankara"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Total payment: ", total)
}
