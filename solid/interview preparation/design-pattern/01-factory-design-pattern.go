package design_pattern

import "fmt"

type Vehicle interface {
	PrintVehicle()
}

type VehicleFactory interface {
	CreateVehicle() Vehicle
}

type Car struct{}

func (c Car) PrintVehicle() {
	fmt.Println("This is a Car")
}

type CarFactory struct{}

func (f CarFactory) CreateVehicle() Vehicle {
	return Car{}
}

type Bike struct{}

func (b Bike) PrintVehicle() {
	fmt.Println("This is a Bike")
}

type BikeFactory struct{}

func (f BikeFactory) CreateVehicle() Vehicle {
	return Bike{}
}

func main() {
	var factory VehicleFactory

	// Use CarFactory
	factory = CarFactory{}
	car := factory.CreateVehicle()
	car.PrintVehicle()

	// Use BikeFactory
	factory = BikeFactory{}
	bike := factory.CreateVehicle()
	bike.PrintVehicle()
}
