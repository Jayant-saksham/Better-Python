//Sorting with Interface
//Create a type that implements Go’s built-in sort.Interface to sort a slice of custom Product structs by price, name, or quantity

package interface_creation_practice

import (
	"fmt"
	"sort"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type SortProcessor interface {
	ISort(product []Product)
}

type SortByPrice struct{}
type SortByName struct{}
type SortByQuantity struct{}

func (s SortByPrice) ISort(product []Product) {
	sort.Slice(product, func(i, j int) bool {
		return product[i].Price < product[j].Price
	})
}

func (s SortByName) ISort(product []Product) {
	sort.Slice(product, func(i, j int) bool {
		return product[i].Name < product[j].Name
	})
}

func (s SortByQuantity) ISort(product []Product) {
	sort.Slice(product, func(i, j int) bool {
		return product[i].Quantity < product[j].Quantity
	})
}

func PrintProducts(products []Product) {
	for _, value := range products {
		fmt.Println("Product Name: ", value.Name)
		fmt.Println("Product Price: ", value.Price)
		fmt.Println("Product Quantity: ", value.Quantity)
	}
}

func main() {
	product1 := Product{
		Name:     "Icecream",
		Price:    20.0,
		Quantity: 10,
	}
	product2 := Product{
		Name:     "Icecream2",
		Price:    30.0,
		Quantity: 90,
	}
	product3 := Product{
		Name:     "Icecream3",
		Price:    10.0,
		Quantity: 2,
	}

	listOfProducts := []Product{product1, product2, product3}

	sortByPrice := SortByQuantity{}
	sortByPrice.ISort(listOfProducts)

	PrintProducts(listOfProducts)
}
