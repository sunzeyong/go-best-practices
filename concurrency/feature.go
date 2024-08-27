package concurrency

import (
	"fmt"
	"time"
)

// pipeline是有前后关系，但是feature是没有前后关系，只需要等待各自做完

func vegatables() <-chan string {
	vega := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vega <- "萝卜 土豆"
	}()
	return vega
}

func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "热水好了"
	}()
	return water
}

func startHotPot() {
	vega := vegatables()
	water := boilWater()
	fmt.Println("sleep a while")
	time.Sleep(2 * time.Second)
	fmt.Printf("等待热水：%v, 蔬菜：%v", <-water, <-vega)
}
