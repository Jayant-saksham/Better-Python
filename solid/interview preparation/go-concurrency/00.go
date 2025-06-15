package main

import "fmt"

func greet() {
	fmt.Println("Hello, World!")
}
func main() {
	go greet()
	fmt.Println("Hello, World in main function!")
}
