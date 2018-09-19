package main

import "fmt"

func main(){
	// 定义两个int类型的channel ch1和ch2
	ch1 := make(chan int,1)
	ch2 := make(chan int,1)
	
	
	// 将1写入ch1
	ch1 <- 1
	
	select{
		case e1 := <- ch1:
			fmt.Printf("1st case is selected. e1=%v",e1)
		
		case e2 := <- ch2:
			fmt.Printf("2ed case is selected. e2=%v",e2)
			
		default:
			fmt.Println("default!")
	}
	// 将ch1写入e1，输出第一个case被选择 e1=...
	// 将ch2写入e2，输出第二个case被选择 e2=...
	// 默认，输出default
	
	// 运行结果：1st case is selected. e1=1
}