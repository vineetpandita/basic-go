package main

import "fmt"

func main() {

	c :=make(chan int)
	done := make(chan bool)


	go func() {
		for i:=10;i<20 ;i++  {
			c <-i
		}
		done <- true
	}()


	go func() {
		for i:=30;i<40 ;i++  {
			c <-i
		}
		done <- true
	}()


	go func() {
		<-done
		<-done
		close(c)
	}()


	for n:= range c {
		fmt.Println(n)
	}

}
