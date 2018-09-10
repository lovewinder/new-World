package main

import "fmt"

func main(){
	array1 := [5]int{1,2,3,4,5}
	var mySlice0 []int = array1[0:5] 
	mySlice0[0] = 24
	fmt.Println(array1,mySlice0) // [24 2 3 4 5] [24 2 3 4 5]
	
	fmt.Println("-----------------------------")
	
	array2 := [5]int{1,2,3,4,5} 
	var mySlice1 []int = array2[0:5]
	mySlice1 = append(mySlice1,24)
	fmt.Println(array2,mySlice1) // [1 2 3 4 5] [1 2 3 4 5 24]
	
	fmt.Println("-----------------------------")
	
	mySlice2 :=  []int{1,2,3,4,5}
	var mySlice3 []int= mySlice2
	mySlice2[0] = 32
	fmt.Println(mySlice2,mySlice3) // [32 2 3 4 5] [32 2 3 4 5]
	
	fmt.Println("-----------------------------")
	
	mySlice4 :=  []int{1,2,3,4,5}
	var mySlice5 []int= mySlice4
	mySlice5 = append(mySlice4,32)
	fmt.Println(mySlice4,mySlice5) // [1 2 3 4 5] [1 2 3 4 5 32]
	
}