package main

import "fmt"

func main(){
	slice:=make([]int ,5,10)
	fmt.Println("len:",len(slice))
	fmt.Println("cap:",cap(slice))
	
	slice2:=[]int{7,8,9}
	
	slice = append(slice,1,2,3)
	fmt.Println("----")
	for _,v:=range slice{
		fmt.Println(v," ")
	}
	
	fmt.Println("----")
	slice = append(slice,slice2...)
	for _,v:=range slice{
		fmt.Println(v," ")
	}
	
	slice3:=[]int{1,2,3,4,5}
	slice4:=[]int{6,7,8}
	
	copy(slice3,slice4)
	fmt.Println("----")
	for _,v:=range slice3{
		fmt.Println(v," ")// 6,7,8,4,5
	}
	
	slice5:=[]int{1,2,3,4,5}
	slice6:=[]int{6,7,8}
	copy(slice6,slice5)
	fmt.Println("----")
	for _,v:=range slice6{
		fmt.Println(v," ")// 1,2,3
	}
	
}