package main

import "fmt"

func f1(){
	fmt.Println("This is f1()")
}

func f2(){
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
	/* This is f1()
	   This is f2()
	   panic:1
	   
	   goroutine 1 [running]:
	main.f2()
		D:/.../.../recover.go:11 +0x8a
	main.main()
		D:/.../.../recover.go:20 +0x2c
	*/
}