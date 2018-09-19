package main

import "fmt"

func main(){
	chanCap := 5
	
	ch := make(chan int,chanCap)
	
	for i := 0;i < chanCap ; i++ {
		select{
		
			case ch <- 1:
		
			case ch <- 2:
		
			case ch <- 3:
		}
	}
	
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch)
	}
}

/* 运行结果：
		第一次：3 3 3 3 2
		第二次：2 3 3 3 3
		第三次：2 2 1 1 2
*/