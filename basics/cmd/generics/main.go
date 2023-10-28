package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("data/contactInfo.json")
	fmt.Printf("\n%v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("data/purchaseInfo.json")
	fmt.Printf("\n%v", purchases)

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	var electircCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println("\n", gasCar, "\n", electircCar)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}
