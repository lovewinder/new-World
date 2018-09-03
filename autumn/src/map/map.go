package main

import "fmt"


	type PersonInfo struct{
		ID string
		Name string
		Address string
	}
func main(){
	// var personDB map[string] PersonInfo// 声明，声明的时候是否分配了内存；没有分配，因为如果跳过创建过程，可以通过编译，但无法给personDB复制
	personDB := make(map[string] PersonInfo)// 创建，如果声明的时候没有分配内存，那么是否可以跳过声明直接创建;如果向跳过声明的话，要使用:=来赋值
	personDB["1234"] = PersonInfo{"1234","y","r2"}
	personDB["1"] = PersonInfo{"1","k","n3"}
	
	person,ok:=personDB["12346"]
	if ok{
		fmt.Println("PersonInfo",person)
	}else{
		fmt.Println("have no person")
	}
}