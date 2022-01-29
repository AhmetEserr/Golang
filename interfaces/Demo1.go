package interfaces

import "fmt"

type Mortgage struct { //kredi
	creditPaymentTotal float64
	address            string
	rate               float64
}
type car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}
func (c car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal

}
func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Istanbul"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 500000, address: "Ankara"}
	credit3 := car{rate: 15, creditPaymentTotal: 600000, carInfo: "Polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("toplam ödeme : ", total)
}
