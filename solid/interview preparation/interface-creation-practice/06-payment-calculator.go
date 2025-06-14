package interface_creation_practice

import "fmt"

type Person struct {
	Name    string
	Bio     string
	Contact string
	Email   string
	Expense int
}

type Team struct {
	Name    string
	persons []Person
	Expense int
}

type Organisation struct {
	Name    string
	teams   []Team
	Expense int
}

type PriceCalculator interface {
	ICalculatePrice() int
}

func (p Person) ICalculatePrice() int {
	return 100
}

func (t Team) ICalculatePrice() int {
	total := 0
	for _, value := range t.persons {
		priceOfIndividual := value.ICalculatePrice()
		total += priceOfIndividual
	}
	return total
}

func (o Organisation) ICalculatePrice() int {
	total := 0
	for _, value := range o.teams {
		priceOfEachTeam := value.ICalculatePrice()
		total += priceOfEachTeam
	}
	return total
}

func main() {
	person1 := Person{
		Name:    "Jayant",
		Bio:     "Bio",
		Contact: "John Doe",
		Email:   "jayant@doe.com",
	}
	person2 := Person{
		Name:    "Sanket",
		Bio:     "Bio",
		Contact: "John Doe",
		Email:   "sanket@doe.com",
	}
	person3 := Person{
		Name:    "Selva",
		Bio:     "Bio",
		Contact: "John Doe",
		Email:   "selva@doe.com",
	}

	team1 := Team{
		Name: "Banking",
		persons: []Person{
			person1, person2, person3,
		},
	}

	team2 := Team{
		Name: "Netbanking",
		persons: []Person{
			person1, person2, person3,
		},
	}

	organisation1 := Organisation{
		Name:  "Razorpay",
		teams: []Team{team1, team2},
	}

	ans := organisation1.ICalculatePrice()
	fmt.Println(ans)

}
