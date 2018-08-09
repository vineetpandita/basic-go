package main

import "fmt"

func main() {

	f1 := simpleFactorial(4);
	fmt.Println(f1)

	//f2 := c_factorial(4)
	//fmt.Println(f2)

}

func simpleFactorial(num int) int {

	if num == 1 {
		return 1
	}
	return num * simpleFactorial(num-1);
}



