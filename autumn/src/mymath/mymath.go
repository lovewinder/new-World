package main

import "fmt"
import "errors"

func Add(a int,b int)(ret int,err error){
	if a<0||b<0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a+b,nil
}

func main(){
	i,err := Add(3,4)
	fmt.Println(i,err) // 7 <nil>
	i2,err2 := Add(-1,2)
	fmt.Println(i2,err2) // 0 Should be non-negative numbers!
}