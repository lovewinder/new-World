package main

import "fmt"

func main(){
	array:=[10]int {1,2,3,4,5,6,7,8,9,10}
	
	slice:=array[:5]
	
	for _,v:= range array{
		fmt.Println(v," ")
	}
	
	fmt.Println("Slice")
	
	for _,v:= range slice{
		fmt.Println(v," ")
	}
	
	mySlice1:=make([]int,5)
	mySlice2:=make([]int,5,10)
	mySlice3:=[]int{1,2,3,4,5}
	
	for _,v:= range mySlice1{
		fmt.Println(v," ")
	}
	for _,v:= range mySlice2{
		fmt.Println(v," ")
	}
	for _,v:= range mySlice3{
		fmt.Println(v," ")
	}
}