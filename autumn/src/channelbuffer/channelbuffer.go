package main

import "fmt"

func main (){
	c := make(chan int,4)
	

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	<- c 	// 先进先出
	c <- 5
	
	// close(c)
	
	for a := range c{
		
		fmt.Println(a)
		fmt.Println("----------")
		
		if a == 4 {
			c <- 11
			close(c)
		}
		
	}
	
}

/* 运行结果：
	2
	----------
	3
	----------
	4
	----------
	5
	----------
	11
	----------
*/
