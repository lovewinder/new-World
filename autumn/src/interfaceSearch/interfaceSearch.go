package main

import (
	"fmt"
	"options"
)

type person struct {
	name string
}

func (p *person) GetName() string {
	return p.name
}

type person2 struct {
	name string
	age int
}

func (p *person2) GetName() string {
	return p.name
}

func (p *person2) GetAge() int {
	return p.age
}

func main() {
	var p options.IPeople = &person{"jack"}
	
	
	// p.(options.IPople2)表示将p转化成options.IPeople2类型，如果成功ok的值为true，否则为false
	if p2,ok := p.(options.IPeople2); ok {
		fmt.Println(p2.GetName(),p2.GetAge())
	} else {
		fmt.Println("p不是IPeople2接口类型") // p不是IPeople2接口类型
	}
	
	var p2 options.IPeople2 = &person2{"mary",23}
	if p,ok := p2.(options.IPeople);ok {
		fmt.Println(p.GetName()) // mary
	}
	
	var pp options.IPeople = &person{"alen"}
	// 将options.IPeople接口类型转化成person指针类型
	if pp2,ok := pp.(*person);ok {
		fmt.Println(pp2.GetName()) // alen
	}
	
	/*
		输出*person，这是接口指向的对象实例的类型。
		但是如果去掉case *person的话，会输出options.IPeople。
	*/
	switch pp.(type){
		/*
		case person:
		 	fmt.Println("person")
		impossible type switch case:pp (type options.IPeople) cannot have dynamic type person (GetName method has pointer receiver)
		*/
		
		case *person:
			fmt.Println("*person")
	
		case options.IPeople2:
			fmt.Println("options.IPeople2")
			
		case options.IPeople:
			fmt.Println("options.IPeople")
			// fallthrough // cannot fallthrough in type switch
			
		
		default:
			fmt.Println("can't found")
	}
	
	var ii interface{} = 43 // 默认int类型
	switch ii.(type) {
		case int:
			fmt.Println("int")
			
		default:
			fmt.Println("can't found")
	}
}