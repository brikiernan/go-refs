package main

import "fmt"

func main() {

	// var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [3]int32{1, 2, 3}
	intArr[1] = 123
	fmt.Println(intArr)
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	intSlice := []int32{4, 5, 6}
	fmt.Printf("This length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThis length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3)
	fmt.Printf("Int slice 3 capacity is %v\n", cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["asdf"])
	var age, ok = myMap2["Adam"]
	fmt.Println(ok)
	// delete(myMap2, "Sarah")
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}
	fmt.Println(myMap2)

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
