package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10) // 定义一个包含10个channel的数组
	for i := 0; i < 10;i++ {
		chs[i] = make(chan int)
		// fmt.Println(chs[i])
		go Count(chs[i]) // 把数组中的每个channel分配给10个不同的goroutine
	}
	
	for _,ch := range(chs) {
		<-ch
	}
	
	chs[1] <- 2
}