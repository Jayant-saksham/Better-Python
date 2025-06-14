package interface_creation_practice

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type RestaurantPriceCalculator interface {
	IGetPrice() float64
}

type Restaurant struct {
	item []Item
}
type EatClub struct {
	priceCalculator []RestaurantPriceCalculator
}

func (r Restaurant) IGetPrice() float64 {
	total := 0.0
	for _, value := range r.item {
		price := float64(value.Price)
		quantity := float64(value.Quantity)
		total += price * quantity
	}
	return total
}

func (e EatClub) IGetPrice() float64 {
	total := 0.0
	for _, value := range e.priceCalculator {
		total += value.IGetPrice()
	}
	return total
}

func main() {
	res1 := Restaurant{
		item: []Item{
			{Name: "test1", Price: 10, Quantity: 10},
			{Name: "test2", Price: 12, Quantity: 50},
			{Name: "test3", Price: 13, Quantity: 60},
			{Name: "test4", Price: 14, Quantity: 70},
		},
	}

	res2 := Restaurant{
		item: []Item{
			{Name: "test6", Price: 15, Quantity: 10},
			{Name: "test7", Price: 0, Quantity: 530},
			{Name: "test9", Price: 3, Quantity: 607},
			{Name: "test411", Price: 1, Quantity: 74},
		},
	}
	e := EatClub{
		priceCalculator: []RestaurantPriceCalculator{
			res1, res2,
		},
	}
	fmt.Println(e.IGetPrice())

}
