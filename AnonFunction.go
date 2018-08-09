package main

import "fmt"


const q int =42



func main() {

	x:=0


	increment := func() int {
		x++
		return x
	}


	fmt.Println(increment())
	fmt.Println(increment())



	a:=43
	b:=67
	c:=787
	d:=789
	fmt.Println(b,c,d)
	fmt.Println(b+c)
	fmt.Println(c+d)
	fmt.Println(d+b)
	fmt.Println(a+b)

	fmt.Println(" Memory  Address for A is ",&a)


	var b1 *int = &a;

	fmt.Println(" Memory  Address for B is ",b)


  	*b1 = *b1 + 1;


	fmt.Println(" Value of A & B is  ",a,*b)



	//All is Pass by Value!



}


