package main

import "fmt"
import "log"

func f1(){
	fmt.Println("This is f1()")
}

func f2(){
	defer func(){
		if err := recover(); err != nil{
			log.Printf("Exception has been caught: %v.",err)
		}
	}()
	
	fmt.Println("This is f2()")
	
	panic(1)
}

func f3(){
	fmt.Println("This is f3()")
}

func main(){
	f1()
	f2()
	f3()
}