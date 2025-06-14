package main

import "fmt"

type Item struct {
	Name     string
	Quantity int
	Price    float64
}

type Orders struct {
	Order map[string]Item
}

func NewOrders() *Orders {
	return &Orders{
		Order: make(map[string]Item),
	}
}

func (o *Orders) AddItems(name string, quantity int, price float64, id string) error {
	itemObject := Item{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
	o.Order[id] = itemObject
	return nil
}

func (o *Orders) TotalPrice() float64 {
	total := 0.0
	for _, value := range o.Order {
		price := value.Price
		quantity := float64(value.Quantity)
		total += price * quantity
	}
	return total
}

type Card struct {
	CardHolderName string
	Cvv            int
	ExpirationDate string
	CardNetwork    string
	CardNumber     string
}

type WalletPaymentPayload struct {
	Email  string
	Wallet int
}

type UPIPaymentPayload struct {
	VPA   string
	Phone string
}

type PaymentProcessor interface {
	IPay(From string, To string)
}

type CardPaymentService struct {
	card Card
}
type WalletPaymentService struct {
	walletPaymentPayload WalletPaymentPayload
}
type UpiPaymentService struct {
	upiPaymentPayload UPIPaymentPayload
}

func (c CardPaymentService) IPay(From string, To string) {
	fmt.Println("Payment processing started from Card Payment Service")
	fmt.Println("From: ", From)
	fmt.Println("To: ", To)
	fmt.Println("Card: ", c.card)
	fmt.Println("CardNetwork: ", c.card.CardNetwork)
	fmt.Println("CardNumber: ", c.card.CardNumber)
	fmt.Println("")
}

func (w WalletPaymentService) IPay(From string, To string) {
	fmt.Println("Payment processing started from Wallet Payment Service")
	fmt.Println("From: ", From)
	fmt.Println("To: ", To)
	fmt.Println("Email: ", w.walletPaymentPayload.Email)
	fmt.Println("Id: ", w.walletPaymentPayload.Wallet)
}

func (u UpiPaymentService) IPay(From string, To string) {
	fmt.Println("Payment processing started from UPI Payment Service")
	fmt.Println("From: ", From)
	fmt.Println("To: ", To)
	fmt.Println("UpiPaymentPayload: ", u.upiPaymentPayload)
}

func main() {
	objectOfOrder := NewOrders()
	err := objectOfOrder.AddItems(
		"Phone",
		2,
		100,
		"123",
	)
	if err != nil {
		fmt.Println(err)
	}

	err2 := objectOfOrder.AddItems(
		"Phone2",
		20,
		60,
		"1234",
	)
	if err2 != nil {
		fmt.Println(err)
	}

	price := objectOfOrder.TotalPrice()
	fmt.Println(price)

	walletPaymentServiceObject := WalletPaymentService{
		walletPaymentPayload: WalletPaymentPayload{
			Email:  "jayant2410@gmail.com",
			Wallet: 10,
		},
	}
	walletPaymentServiceObject.IPay("Jayant", "Urvashi")
}
