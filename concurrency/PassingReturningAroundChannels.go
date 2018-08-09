package main

import (
	"fmt"
)

func main() {

	c := incremntor()
	s := adder(c)
	for n := range s {
		fmt.Println(" GoRoutine",n)
	}


	var sum int
	//go func() {
		for i:=1; i<10000000;i++ {
			sum = sum + i
		}
		fmt.Println(" NoGo Route",sum)
	//}()
}


func incremntor() chan int {
	out :=make(chan int)
	go func() {
		for i:=1; i<10000000;i++ {
			out<-i
		}
		close(out)
	}()
	return out
}


func adder (c <-chan int) chan int {
	out :=make(chan int)
	go func() {
		var sum int
		for n:=range c {
			sum = sum + n
		}
		out <- sum
		close(out)
	}()
	return out
}