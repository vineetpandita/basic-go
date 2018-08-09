package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"runtime"
)


var wg sync.WaitGroup

var mutex sync.Mutex


var counter int

func init ()  {

	runtime.GOMAXPROCS(runtime.NumCPU())

}


func main() {

/*	wg.Add  (2)

	go foo()
	go bar()

	wg.Wait()*/

	wg.Add(2)

	go incr1(" T1 ")
	go incr1(" T2 ")


	wg.Wait()

fmt.Println("Final COunter ", counter)


}
func incr1(s string) {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<100; i++ {

		time.Sleep(time.Duration(rand.Intn(20))*  time.Millisecond)

		//mutex.Lock()
		counter++
 		//mutex.Unlock()

		fmt.Println(s," - ",counter, " @ iteration ", i)
	}

	wg.Done()
}

func foo ()  {

	for i:=1; i<1000; i++ {
		fmt.Println(" Foo ", i)
	}

	wg.Done()
}

func bar()  {


	for i:=1; i<1000; i++ {
		fmt.Println(" Bar ", i)
	}

	wg.Done()
}


