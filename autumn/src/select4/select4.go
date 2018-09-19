package main

import "fmt"


func main(){
	var ch chan int = make(chan int,1)
	ch <- 1
	
	select {
		case <- ch:
			fmt.Println("before")
			break
			fmt.Println("after")
		default:
			fmt.Println("default")
	}
	
	fmt.Println("over")
}

/* 运行结果：
		before
		over
*/