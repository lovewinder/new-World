package main

// go的函数不存在重载，因此不可以有重名函数
// go可以根据“类型”的函数，自动生成“类型指针”的函数，但是无法根据“类型指针”的函数生成“类型”函数。
import "fmt"

type Integer int

/*
	类型Integer实现的两个方法是类型调用的Less方法和类型指针调用的Add方法，
	当你用类型实例给接口赋值的时候，具体对应两个方法时发现
	func (a Integer) Less(b Integer) bool {
		return a < b
	}
	可以对应上，但是
	func (a *Integer) Add(b Integer) {
		*a += b 
	}
	这个方法需要接收的是一个指针，但现在只传进来一个实例，所以编译失败。
	当你用类型指针实例给接口赋值的时候，具体对应两个方法时发现
	func (a Integer) Less(b Integer) bool {
		return a < b
	}
	无法对应上，但是“类型”的函数可以自动生成“类型指针”函数，所以通过了。
	func (a *Integer) Add(b Integer) {
		*a += b 
	}
	这个方法本身就可以对应上。
*/

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) LessAuto(b Integer) bool {
	return (*a).Less(b)
}

func (a *Integer) Add(b Integer) {
	*a += b 
}

func (a Integer) AddAutoFail(b Integer) {
	(&a).Add(b)
}

type Integer2 int

func (a Integer2) Less(b Integer2) bool {
	return a < b
}

func (a Integer2) Add(b Integer2) {
	a += b 
}

type LessAdd interface {
	Less(b Integer) bool
	Add(b Integer)
}

type LessAdd2 interface {
	Less(b Integer2) bool
	Add(b Integer2)
}

func main() {
	var a Integer = 7
	fmt.Println("print1",a) // print1 7
	a.AddAutoFail(1)
	fmt.Println("print2","AddAutoFail",a) // print2 AddAutoFail 7
	a.Add(1)
	fmt.Println("print2","Add",a) // print2 Add 8
	// var b LessAdd = a // cannot a (type Integer) as type LessAdd in assignment: Integer does not implement LessAdd (Add method has pointer receiver)
	// fmt.Println("print3","Show",b)
	var b LessAdd = &a
	fmt.Println("print3","Show",b) // print3 Show 0x04200c0c0
	var c Integer2 = 21
	var d LessAdd2 = c
	fmt.Println("print4","Show2",d) // print4 Show2 21
}