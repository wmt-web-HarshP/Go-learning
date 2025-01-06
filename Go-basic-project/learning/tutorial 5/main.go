package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = []rune("résumé") 
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed) 
	for i, v := range myString{
		fmt.Println(i,v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len (myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"} 
	// var catStr = ""
	var strBuilder strings.Builder
	for i := range strSlice{
		// catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}
	var catStr=strBuilder.String()
	fmt.Printf("\n%v", catStr)
}