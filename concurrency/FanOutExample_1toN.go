package main

import "fmt"

func main() {


	c:= make( chan int)
	done :=make(chan bool)



	go func() {
		for i:=0; i<10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println("A-",n)
		}
		done <- true
	}()


	go func() {
		for n := range c {
			fmt.Println("B-",n)
		}
		done <- true
	}()

	//Sort of waiting barrier for above threads to complete!
	<-done
	<-done

}
