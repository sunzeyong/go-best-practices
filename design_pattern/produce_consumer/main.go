package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 10)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go producer(wg, ch)
	go consumer(wg, ch)
	wg.Wait()
}

func producer(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for {
		productId := strconv.Itoa(time.Now().Nanosecond())
		ch <- "商品ID" + productId
		fmt.Println("生产了商品" +  productId)
		time.Sleep(time.Second)
	}
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for {
		product := <-ch
		fmt.Println("消费了商品"+ product)
		time.Sleep(time.Second)
	}
}


