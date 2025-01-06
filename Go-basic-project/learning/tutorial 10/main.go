package main

import "fmt"

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
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	var electricCars = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   124,
			mpkwh: 40,
		},
	}
	fmt.Println(gasCar, electricCars)
}

// func main() {
// 	var intSlice = []int{1, 2, 3, 4}
// 	fmt.Println(sumOfArr(intSlice))

// 	var float32Slice = []float32{2, 3, 4, 5}
// 	fmt.Println(sumOfArr(float32Slice))
// }

// func sumOfArr[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }
