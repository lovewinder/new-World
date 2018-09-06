package main

import(
	"fmt"
)

func adder() func(int) int{
		sum := 0
		return func(x int) int{
			sum += x
			return sum
		}
}

func main(){
	var j int = 5
	a := func()(func()){
		var i int = 10
		return func(){
			fmt.Printf("i,j:%d,%d\n",i,j)
		}
	}()
	// i *= 2 // undefined i
	a() // i,j:10,5
	j *= 2
	a() // i,j:10,10
	
	
	pos,neg := adder(),adder()
	for i := 0;i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i), // 这里的价值就体现在闭包中的sum值不会因为别的函数调用而改变。如果没有闭包，那么sum := 0应该设置在哪里合适？如果作为全局变量，在pos()调用后，sum的值就会改变，需要在每次调用pos()或neg()函数之后，添加sum = 0语句。如果作为函数内的局部变量，在没有闭包的情况下，adder()函数中想完成同样的操作需要写for循环，在循环中sum总是被赋值为0，不符合实际情况。
		)
	}
	
}

