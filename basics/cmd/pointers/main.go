package main

import "fmt"

func main() {
	var pointer *int32 = new(int32)
	var int int32

	fmt.Println(pointer, int)
	fmt.Printf("\nThe value p points to is: %v", *pointer)
	fmt.Printf("\nThe value of i is: %v", int)
	*pointer = 10
	fmt.Printf("\nThe value p points to is: %v", *pointer)
	pointer = &int
	*pointer = 30
	fmt.Printf("\nThe value p points to is: %v", *pointer)
	fmt.Printf("\nThe value of i is: %v", int)
	var k int32 = 40
	int = k
	fmt.Printf("\nThe value of i is: %v\n", int)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4

	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("The thing1 is: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of thing2 array is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
