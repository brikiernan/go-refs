package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	fmt.Println(len("🤷‍♂️"))
	fmt.Println(utf8.RuneCountInString("🤷‍♂️"))
	fmt.Println(len("aasdasd"))
	var myRune rune = 'a'
	fmt.Println(myRune)

	var randomBoolean bool = rand.Float32() >= 0.5
	fmt.Println(randomBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pie float32 = 3.1415926
}
