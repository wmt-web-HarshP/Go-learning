package main

import (
	"fmt"
	"time"
)

func main(){
	var n int =1000000;
	var testSlice=[]int{}
	var testSlice2=make([]int,0,n)	

	fmt.Println("Total time without preallocation:%v\n", timeLoop(testSlice,n))
	fmt.Println("Total time with preallocation:%v",timeLoop(testSlice2,n))
	fmt.Println("Hello WOrld")
}

func timeLoop(slice []int, n int)time.Duration{
	var t0=time.Now()
	for len(slice)<n{
		slice=append(slice, 1)
	}
	return time.Since(t0)
}

// package main

// import "fmt"

// func main() {

// 	//array
// 	// var intArr []int32 = []int32{1, 2, 3, 4, 5, 6}
// 	intArr := [...]int32{1, 2, 3, 4, 5}
// 	fmt.Println(intArr)

// 	var intSlice []int32 = []int32{4, 5, 6}
// 	fmt.Println(intSlice)
// 	intSlice = append(intSlice, 7)
// 	fmt.Println(intSlice)

// 	var intSlice1 []int32 = []int32{8, 9}
// 	intSlice = append(intSlice, intSlice1...)
// 	fmt.Println(intSlice)

// 	var intSlice2 []int32 = make([]int32, 3, 10)
// 	fmt.Println(intSlice2)

// 	var myMap map[string]uint8 = make(map[string]uint8)
// 	fmt.Println(myMap)

// 	var myMap2 = map[string]uint8{"Harsh": 25, "vatsal": 32}
// 	fmt.Println(myMap2["Harsh"])
// 	// intArr[1]=123
// 	// fmt.Println(intArr[0])
// 	// fmt.Println(intArr[1:3])

// 	// fmt.Println(&intArr[0])
// 	// fmt.Println(&intArr[1])
// 	// fmt.Println(&intArr[2])

// 	//if-else
// 	var age, ok = myMap2["sagar"]
// 	if ok {
// 		fmt.Println("The age is %v", age)
// 	} else {
// 		fmt.Println("Invalid Name")
// 	}
// 	//for-loop
// 	for name, age := range myMap2 {
// 		fmt.Println("Name:%v,Age:%v \n", name, age)
// 	}

// 	for i, v := range intArr {
// 		fmt.Println("Index:%v,Value:%v \n", i, v)
// 	}

// 	var i int = 0
// 	for {
// 		if i >= 10 {
// 			break
// 		}
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for i = 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }
