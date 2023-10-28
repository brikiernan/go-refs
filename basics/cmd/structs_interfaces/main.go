package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	myEngine.gallons = 10
	fmt.Println(myEngine, myEngine.name)

	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
		owner
	}{35, 12, owner{"Chase"}}

	fmt.Println(myEngine2, myEngine2.name)

	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 50)

	var myEEngine electricEngine = electricEngine{25, 15}

	fmt.Printf("Total miles left in tank: %v\n", myEEngine.milesLeft())
	canMakeIt(myEEngine, 50)

}
