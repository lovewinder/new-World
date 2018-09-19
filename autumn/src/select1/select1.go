package main

import "fmt"
import "time"


	func f1(ch chan int){
		time.Sleep(time.Second * 5)
		ch <- 1
	}
	
	func f2(ch chan int){
		time.Sleep(time.Second *10)
		ch <- 1
	}
	
	func main(){
		ch1 := make(chan int,1)
		ch2 := make(chan int,1)
		
		go f1(ch1)
		go f2(ch2)
		
		select{
			case <- ch1 :
				fmt.Println("1 case")
			
			case <- ch2 :
				fmt.Println("2 case")
		}
	
	// 运行结果：等待5秒后，"1 case"