package main

import (
	"fmt"
	"runtime"
)


func init ()  {

	runtime.GOMAXPROCS(1)

}

func main() {


	c :=make( chan int)


	go func() {
		for i:=0; i<10 ;i++  {
			c <- i

			fmt.Println("Putting In ",i)

		}
		close(c)
	}()


/*	go func() {
		for {
			fmt.Println(<-c)
		}
	}()*/


	for n:=range c{
		fmt.Println("%t",n)
	}


	//time.Sleep(time.Second)
}
