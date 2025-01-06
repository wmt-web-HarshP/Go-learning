package main

import "fmt"

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\n The memory location of the thing2 array is:%p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

//pointer are special type of variables
func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n The memory location of the thing1 array is:%p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is:%v", result)
	fmt.Printf("\nThe value of thing1 is:%v", thing1)

	//Take care thing
	// var arr=[]int32{1,2,3}
	// var arrCopy=arr
	// arrCopy[2]=4 //it will also change the original array value
	// fmt.Println(arr)
	// fmt.Println(arrCopy)

	// var p*int32=new(int32)
	// var i int32
	// fmt.Println("The value p points to is:%v",*p)
	// fmt.Println("\n The value if i is %v",i)
	// p=&i
	// *p=1
	// fmt.Println("The value p points to is:%v",*p)
	// fmt.Println("\n The value if i is %v",i)
	// var k int32=2
	// i=k
	// fmt.Println("\n The value if i is %v",i)
	// fmt.Println("\n The value if k is %v",k)
}
