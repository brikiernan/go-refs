package main

import (
	"fmt"
	"strings"
)

func main() {

	// var myString = "résumé"
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T \n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v\n", len(myString))

	var myRune = '🚀'
	fmt.Printf("\nmyRune = %v as type %T\n", myRune, myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	fmt.Println(strSlice)
	var strBuilder strings.Builder
	var catStr = ""
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // efficient, much faster, appends value i
		catStr += strSlice[i]               // inefficient, new string each iteration
	}
	var builtStr = strBuilder.String()
	fmt.Printf("%v\n", builtStr)
	fmt.Printf("%v\n", catStr)

}
