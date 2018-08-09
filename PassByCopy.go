package main

import "fmt"

func zero( x int, y *int)  {

	x = 0;

	*y = 0;

}


func main() {


	x :=5
	y :=5

	zero(x,&y)



	fmt.Println("x ",x)
	fmt.Println("y",y)




}