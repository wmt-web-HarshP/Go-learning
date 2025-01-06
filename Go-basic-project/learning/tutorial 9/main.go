package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PANNER_PRICE float32 = 5
var MAX_CHEEZE_PRICE float32 = 3

func main() {
	var PANNERCHANNEL = make(chan string)
	var CHEEZECHANNEL = make(chan string)
	var website = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range website {
		go checkPannerPrices(website[i], PANNERCHANNEL)
		go checkCheezePrice(website[i], CHEEZECHANNEL)
	}
	sendMessage(PANNERCHANNEL, CHEEZECHANNEL)
}

func checkCheezePrice(website string, c chan string) {
	for {
		time.Sleep(time.Second*1)
		var cheeze_price=rand.Float32()*20
		if cheeze_price<MAX_CHEEZE_PRICE{
			c<-website
			break
		}
	}
}

func checkPannerPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second*1)
		var panner_price=rand.Float32()*20
		if panner_price<MAX_PANNER_PRICE{
			c<-website
			break
		}
	}
}

func sendMessage(PANNERCHANNEL chan string,CHEEZECHANNEL chan string){
	select{
	case website:=<-PANNERCHANNEL:
		fmt.Printf("\nText sent:Found deal on panner at %v",website)
		
	case website:=<-CHEEZECHANNEL:
		fmt.Printf("\nText sent:Found deal on cheeze at %v",website)
	}
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var c = make(chan int, 5)
// 	go process(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func process(c chan int) {
// 	defer close(c)
// 	for i:= 0; i < 5; i++ {
// 		c <- 1
// 	}
// 	fmt.Println("Exiting process")
// }
