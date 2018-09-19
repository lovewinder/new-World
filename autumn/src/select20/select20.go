package main

import "fmt"


var ch1 chan int

var ch2 chan int

var chs = []chan int{ch1,ch2} // ch1和ch2都是空值channel，和空值channel通信永远也不会成功。

var numbers = []int{1,2,3,4,5}

func main() {

	select {
		case getChan(0) <- getNumber(2):
			fmt.Println("1 case")
			
		case getChan(1) <- getNumber(3):
			fmt.Println("2 case")
			
		default:
			fmt.Println("default!")
	}
}

func getNumber(i int) int{
	fmt.Printf("numbers[%d]\n",i)
	return numbers[i]
}

func getChan(i int) chan int{
	fmt.Printf("chs[%d]\n",i)
	return chs[i]
}
	
/* 运行结果：
	chs[0]
	numbers[2]
	chs[1]
	numbers[3]
	default!
*/